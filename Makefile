IMAGE	:= cirocosta/go-tcp-echo
SRCS	:= $(shell find . -name "*.go" -type f -maxdepth 1)

all: build

format: $(SRCS)
	gofmt -s -w .

build: format
	go build -v .

image:
	docker build -t $(IMAGE) .

.PHONY: image build format

