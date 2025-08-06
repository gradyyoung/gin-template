export GOOS=linux &&\
export GOARCH=amd64 &&\

cd ./cmd/wire &&\
wire &&\

cd ../../ &&\
go build -o gin-template ./cmd/main.go