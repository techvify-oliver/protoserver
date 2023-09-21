# --go_out va --go-grpc_out shoud be the same

GO_MODULE=github.com/techvify-oliver/protoserver

gen_contract-user:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/contract_user/*.proto

gen_auth-user:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/auth_user/*.proto