kitex:
	@rm -rf kitex_gen/
	@kitex -module github.com/zlllgp/vegas -service vegas ./idl/vegas.proto

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