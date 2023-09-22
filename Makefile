# --go_out va --go-grpc_out shoud be the same

GO_MODULE=github.com/techvify-oliver/protoserver

.PHONY: git tag

gen:
	protoc --proto_path=./proto --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/*.proto

git:
	.\gitcombo.bat

tag: # get latest tag
	git describe --tags $(git rev-list --tags --max-count=1)