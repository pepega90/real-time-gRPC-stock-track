web:
	cd ./http && go run main.go

server:
	cd ./grpc && go run main.go

c:
	cd ./client && go run main.go