RESOURCE_PROTO_FILES= $(shell find api/resources -type f -name '*.proto')
SERVICE_PROTO_FILES= $(shell find api/services -type f -name '*.proto')

.PHONY: build
build: deps
	protoc \
	    --go_opt=module=github.com/dstreamcloud/proto \
	    --go-grpc_opt=module=github.com/dstreamcloud/proto \
	    --proto_path=./api \
	    --go-grpc_out=./ \
	    --go_out=./ \
	    $(RESOURCE_PROTO_FILES) $(SERVICE_PROTO_FILES)
	protoc \
		--clone_opt=module=github.com/dstreamcloud/proto \
		--setter_opt=module=github.com/dstreamcloud/proto \
		--proto_path=./api \
	    --clone_out=./ \
	    --setter_out=./ \
		$(RESOURCE_PROTO_FILES)


.PHONY: deps
deps:
	go install github.com/joesonw/proto-tools/cmd/protoc-gen-clone
	go install github.com/joesonw/proto-tools/cmd/protoc-gen-setter

