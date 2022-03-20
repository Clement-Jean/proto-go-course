build: generate
	go build .

generate:
	protoc -Iproto --go_opt=module=github.com/Clement-Jean/proto-go-course --go_out=. proto/*.proto

clean:
	rm proto/*.pb.go
	rm proto-go-course