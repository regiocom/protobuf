regenerate:
	go install github.com/regiocom/protobuf/protoc-min-version
	go install github.com/regiocom/protobuf/protoc-gen-gogoslick
	protoc-min-version --version="3.0.0" --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/regiocom/protobuf/protobuf/ \
	--gogoslick_out=Mgoogle/protobuf/timestamp.proto=github.com/regiocom/protobuf/types:. timefail.proto
