gen:
	protoc --version
	protoc -I=. --go_opt=paths=source_relative  --go_out=plugins=grpc:.  ./api/**/*.proto
	go mod tidy

clean:
	rm -rf ./api/**/*.pb.go
	go mod tidy