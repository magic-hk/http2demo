# Option #
# ############################################
# For need to create new cert.pem and key.pem
# server Cert and Key generate
cd pem
go run generate_cert.go --host localhost
# ###########################################


# Download lib #
# ###########################################
# Before RUN 
# set mod env
# open this project in vs code
# ctrl + ?to get Terminal 

## set go mod repository
# Linux or macOs
export GOPROXY=https://goproxy.io

# Windows
$env:GOPROXY="https://goproxy.io"


# get lib 
go mod download
# ###########################################



# How ro run #
# ###########################################
cd server
go run ./server.go

cd client
go run ./client.go
# ###########################################
