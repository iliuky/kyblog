package middleware

import (
	"bytes"
	"io/ioutil"
	"kyblog/internal/common"
	"kyblog/internal/tools/xos"
	"os"
	"path"
	"time"

	"kyblog/internal/blog/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	bodyBuf *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	//memory copy here!
	w.bodyBuf.Write(b)
	return w.ResponseWriter.Write(b)
}

// PageStatic 页面静态化
func PageStatic(s *service.Service) gin.HandlerFunc {
	// map 中存在才能进行访问, 否则返回404, 等于true 表示已经静态化
	pageMap := buildPageMap(s)
	go func() {
		for {
			// 定时判断文章是否有更新
			pageMap = updatePageMap(s, pageMap)
			time.Sleep(time.Second * 10)
		}
	}()
	pwd, _ := os.Getwd()
	os.MkdirAll(path.Join(pwd, "./cache/a"), os.ModePerm)

	return func(c *gin.Context) {
		uri, fullPath := getStaticFilePath(c)
		exists, ok := pageMap[uri]
		if !ok {
			c.Status(404)
			c.Abort()
			return
		}

		// 如果存在静态文件
		filePath := path.Join(pwd, "./cache/"+uri+".html")
		if exists {
			c.File(filePath)
			c.Abort()
			return
		}

		blw := bodyLogWriter{bodyBuf: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		c.Next()
		if c.Writer.Status() == 200 {
			err := ioutil.WriteFile(filePath, blw.bodyBuf.Bytes(), 0644)
			if err != nil {
				common.Logger.Error(err.Error(), zap.String("filePath", filePath))
				return
			}

			pageMap[uri] = true
			if fullPath == "/a/:pinyin" {
				pinyin := c.Param("pinyin")
				s.UpdateArticleStaticTime(pinyin)
			}
		}
	}
}

func buildPageMap(s *service.Service) (_pageMap map[string]bool) {
	var keys = []string{}
	keys = append(keys, "index")
	articles := s.GetArticleAll()
	for _, a := range articles {
		keys = append(keys, "a/"+a.PinYin)
	}

	pwd, _ := os.Getwd()
	mp := make(map[string]bool)
	for _, val := range keys {
		filePath := path.Join(pwd, "cache", val+".html")
		isExists, _ := xos.Exists(filePath)
		mp[val] = isExists
	}

	return updatePageMap(s, mp)
}

func updatePageMap(s *service.Service, _pageMap map[string]bool) map[string]bool {
	mp := make(map[string]bool)
	for k, v := range _pageMap {
		mp[k] = v
	}

	statics := s.GetArticleNonStaticAll()
	isChanage := false
	for _, py := range statics {
		key := "a/" + py.PinYin
		v, ok := mp[key]
		if !ok || v {
			mp[key] = false
			isChanage = true
		}
	}

	if isChanage {
		mp["index"] = false
		return mp
	}
	return _pageMap
}

func getStaticFilePath(c *gin.Context) (uri string, fullPath string) {
	fullPath = c.FullPath()
	if fullPath == "/" {
		uri = "index"
	}
	if fullPath == "/a/:pinyin" {
		uri = "a/" + c.Param("pinyin")
	}
	return
}
