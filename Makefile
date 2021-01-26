RESOURCE_PROTO_FILES= $(shell find api/resources -type f -name '*.proto')
SERVICE_PROTO_FILES= $(shell find api/services -type f -name '*.proto')

.PHONY: default
default: build

.PHONY: build
build: deps
	protoc \
	    --go_opt=module=github.com/dstreamcloud/dt-proto \
	    --go-grpc_opt=module=github.com/dstreamcloud/dt-proto \
	    --proto_path=./api \
	    --go-grpc_out=./ \
	    --go_out=./ \
	    $(RESOURCE_PROTO_FILES) $(SERVICE_PROTO_FILES)
	protoc \
		--clone_opt=module=github.com/dstreamcloud/dt-proto \
		--setter_opt=module=github.com/dstreamcloud/dt-proto \
		--proto_path=./api \
	    --clone_out=./ \
	    --setter_out=./ \
		$(RESOURCE_PROTO_FILES)


.PHONY: deps
deps:
	go install github.com/joesonw/proto-tools/cmd/protoc-gen-clone
	go install github.com/joesonw/proto-tools/cmd/protoc-gen-setter
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc

