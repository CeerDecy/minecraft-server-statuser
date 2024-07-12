TAG="1.0.0"

build:
	docker buildx build --platform linux/amd64,linux/arm64 -t ceerdecy/minecraft-server-statuser:${TAG} .

build-push:
	docker buildx build --platform linux/amd64,linux/arm64 -t ceerdecy/minecraft-server-statuser:${TAG} --push .
