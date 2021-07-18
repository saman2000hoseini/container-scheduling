export APP=container-scheduling
export LDFLAGS="-w -s"

run:
	go run -ldflags $(LDFLAGS) ./cmd/container-scheduling server

build:
	go build -ldflags $(LDFLAGS) ./cmd/container-scheduling

install:
	go install -ldflags $(LDFLAGS) ./cmd/container-scheduling

build-image:
	docker build -t container_scheduling ./assets/container
