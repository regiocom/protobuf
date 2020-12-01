regenerate:
	go install github.com/regiocom/protobuf/protoc-gen-gogo
	protoc --proto_path=.:$(GOPATH)/src/:$(GOPATH)/src/github.com/regiocom/protobuf/protobuf/ \
	--gogo_out=. events.proto 
