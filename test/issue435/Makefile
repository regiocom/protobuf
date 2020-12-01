regenerate:
	go install github.com/regiocom/protobuf/protoc-gen-gogo
	protoc --gogo_out=. --proto_path=../../../../../:../../protobuf/:. *.proto
