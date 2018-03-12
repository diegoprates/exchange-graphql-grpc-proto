protoPath := $(dir $(abspath $(lastword $(MAKEFILE_LIST))))

.PHONY: proto

proto: 
	docker run --rm -v $(protoPath):$(protoPath) -w $(protoPath) protoc_graphql_builder:latest -I=$(protoPath) --gogoopsee_out=plugins=graphql+micro,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. --micro_out=. $(protoPath)$(service)/$(service).proto