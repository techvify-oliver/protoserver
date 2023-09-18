# --go_out va --go-grpc_out phai giong nhau

GO_MODULE=github.com/techvify-oliver/grpcmicro

gen_greet:
	protoc --go_opt=module=${GO_MODULE} --go_out=. \
	--go-grpc_opt=module=${GO_MODULE} --go-grpc_out=. \
	./proto/*.proto

greet_client:
	go build learn_grpc_set1/greet/client
	./client

greet_server:
	go build learn_grpc_set1/greet/server
	./server

gen_calculator:
	protoc --go_out=./calculator/generated --go-grpc_out=./calculator/generated calculator/proto/*.proto

calculator_client:
	go build learn_grpc_set1/calculator/client
	./client

calculator_server:
	go build learn_grpc_set1/calculator/server
	./server