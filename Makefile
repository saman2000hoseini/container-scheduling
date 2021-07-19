export APP=container-scheduling
export LDFLAGS="-w -s"

run:
	go run -ldflags $(LDFLAGS) ./cmd/container-scheduling cli

run-server:
	go run -ldflags $(LDFLAGS) ./cmd/container-scheduling server

build:
	go build -ldflags $(LDFLAGS) ./cmd/container-scheduling

install:
	go install -ldflags $(LDFLAGS) ./cmd/container-scheduling

build-image:
	docker build -t container_scheduling ./assets/container

run-containers:
	docker run -d -t --name container_scheduler1 container_scheduling
	docker run -d -t --name container_scheduler2 container_scheduling
	docker run -d -t --name container_scheduler3 container_scheduling

docker-cleanup:
	docker stop container_scheduler1 container_scheduler2 container_scheduler3
	docker rm container_scheduler1 container_scheduler2 container_scheduler3
	docker rmi container_scheduling
