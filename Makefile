# --go_out va --go-grpc_out shoud be the same

GO_MODULE=github.com/techvify-oliver/protoserver


.PHONY: git


gen_contract-user:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/contract_user/*.proto

gen_contract-group-request:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/contract_request_group/*.proto

git:
	.\gitcombo.bat