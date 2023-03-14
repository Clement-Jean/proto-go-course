# proto-go-course

[![build main branch](https://github.com/Clement-Jean/proto-go-course/actions/workflows/build.yml/badge.svg)](https://github.com/Clement-Jean/proto-go-course/actions/workflows/build.yml) [![Lint protobuf](https://github.com/Clement-Jean/proto-go-course/actions/workflows/lint.yml/badge.svg)](https://github.com/Clement-Jean/proto-go-course/actions/workflows/lint.yml)

## COUPON: `START_MAR`

## Notes

### `Windows`

- I recommend you use powershell (try to update: [see](https://github.com/PowerShell/PowerShell/releases)) for following this course.
- I recommend you use [Chocolatey](https://chocolatey.org/) as package installer (see [Install](https://chocolatey.org/install))


### Build

#### `Prerequisites`

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

#### `Linux/MacOS`

```shell
go mod tidy
make
```

#### `Windows - Chocolatey`
```shell
choco install make
go mod tidy
make
```

#### `Windows - Without Chocolatey`

```shell
protoc -Iproto --go_opt=module=github.com/Clement-Jean/proto-go-course --go_out=. proto/*.proto

go mod tidy
go build -o proto-go-course.exe .
```

## Run

```
./proto-go-course
```

or

```
./proto-go-course.exe
```
