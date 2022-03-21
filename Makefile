PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')

build: 	generate
	go build .

generate:
	protoc -Iproto --go_opt=module=${PACKAGE} --go_out=. proto/*.proto

clean:
	rm proto/*.pb.go
	rm proto-go-course