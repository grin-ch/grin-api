protoc --version
protoc -I=. --go_opt=paths=source_relative  --go_out=plugins=grpc:.  ./api/grpc/**/*.proto
go mod tidy