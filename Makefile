PROJECT_NAME = "terminal-dance"
DOCKER_IMAGE_NAME = "tasnimzotder/$(PROJECT_NAME)"

docker-build:
	docker build -t $(DOCKER_IMAGE_NAME) .

docker-run:
	docker run -it --rm -v $(PWD)/frames:/app/frames $(DOCKER_IMAGE_NAME) sh -c "make local-build && ./main $(size)"

local-build:
	go build -o main . && chmod +x main

local-run:
	./main $(size)

.PHONY: docker-build docker-run local-build local-run