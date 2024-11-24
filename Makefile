gen:
	rm -rf kitex_gen/
	kitex -no-fast-api -module github.com/zlllgp/vegas ./api/errno.proto
	kitex -no-fast-api -module github.com/zlllgp/vegas ./api/base.proto
	kitex -no-fast-api -service vegas -module github.com/zlllgp/vegas -I ./api ./api/vegas.proto
	kitex -no-fast-api -service right -module github.com/zlllgp/vegas -I ./api ./api/right.proto

# 定义变量
IMAGE_NAME = vegas
CONTAINER_NAME = vegas
APP_NAME = vegas

# 构建Docker镜像 --progress=plain
build:
	docker build -t ${IMAGE_NAME} -f Dockerfile --progress=plain .

# 运行Docker容器
run:
	docker run -d --name $(CONTAINER_NAME) -p 8080:8080 $(IMAGE_NAME)

# 停止Docker容器
stop:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)

# 清理Docker容器和镜像
clean:
	docker stop $(CONTAINER_NAME)
	docker rm $(CONTAINER_NAME)
	docker rmi $(IMAGE_NAME)