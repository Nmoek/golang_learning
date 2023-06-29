#!/bin/bash

SRC_PATH="./"
BIN_PATH=${SRC_PATH}/../



SERVER_SRC=(
    ${SRC_PATH}/1_文件传输服务器.go
    ${SRC_PATH}/msgProtocol.go
)

CLIENT_SRC=(
    ${SRC_PATH}/2_文件传输客户端.go
    ${SRC_PATH}/msgProtocol.go
)

rm -rf ${SRC_PATH}/server ${SRC_PATH}/client

go build -o server.exe ${SERVER_SRC[*]}
go build -o client.exe ${CLIENT_SRC[*]}
