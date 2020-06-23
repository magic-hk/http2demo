# Option #
## For need to create new cert.pem and key.pem 
```
# server Cert and Key generate
cd pem
go run generate_cert.go --host localhost
```



# Download lib #
```
# cd project dir
cd /opt/workspace/gomod/http2demo

# set go mod repository
# Linux or macOs
export GOPROXY=https://goproxy.io

# Windows
$env:GOPROXY="https://goproxy.io"

# Download lib 
go mod download
```

# How ro run #
```
# run server
cd /opt/workspace/gomod/http2demo/server
go run ./server.go

# run client
cd /opt/workspace/gomod/http2demo/client
go run ./client.go
```