# 停止容器和删除镜像
docker stop blog;
docker rm blog;
docker rmi blog:1.0

# 构建镜像运行镜像
docker build -t blog:1.0 . ;
docker run -d -p 8020:80 --restart always --name blog \
--env ENVIRONMENT=production \
--add-host dbhost:192.168.0.21 \
blog:1.0

