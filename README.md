# grpc-rest-gateway
This repo implements a grpc micro service having the grpc calls exposed as Rest calls  through a rest gateway
TODO: Rest Gateway

# Install protoc
```
Download the protobuf-<latest version> from https://github.com/protocolbuffers/protobuf/releases/
Extract the tar.gz file.
$cd ~/Downloads/protobuf-2.4.1
$./configure
$make
$make check
$sudo make install
$which protoc
$protoc --version
```

# Download and install protoc-gen-go
```
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

