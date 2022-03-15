gen:
	protoc --version
	protoc --gogo_out=plugins=grpc:. ./api/**/*.proto
	go mod tidy

clean:
	rm -rf ./api/**/*.pb.go
	go mod tidy