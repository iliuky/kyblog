package middleware

import (
	"bytes"
	"io/ioutil"
	"kyblog/internal/common"

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

// map 中存在才能进行访问, 否则返回404, 等于true 表示已经静态化
var pageMap map[string]bool = make(map[string]bool)

// PageStatic 页面静态化
func PageStatic() gin.HandlerFunc {
	pageMap["index"] = false
	return func(c *gin.Context) {
		uri := c.FullPath()
		if uri == "/" {
			uri = "index"
		}
		exists, ok := pageMap[uri]
		if !ok {
			c.Status(404)
			c.Abort()
			return
		}

		// 如果存在静态文件
		filePath := "./cache/" + uri + ".html"
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
		}
	}
}
