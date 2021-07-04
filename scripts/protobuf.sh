#!/bin/bash

ROOT_PATH=/go/src/github.com/jamiri/port-app
GO_OUTPUT=/go/src

echo "Generating proto messages...";
APPS=$(ls ${ROOT_PATH}/protocol/message)
for APP in ${APPS}; do
    FILE_PATH=${ROOT_PATH}/protocol/message/${APP}/*.proto
	protoc -I/protobuf \
		-I${ROOT_PATH} \
		--go_out=plugins=grpc:${GO_OUTPUT} \
		${FILE_PATH}
    echo "[${APP}]";
done

echo "Generating services...";
SRV=$(ls ${ROOT_PATH}/protocol/service)
for SRV in $SRV; do
    FILE_PATH=${ROOT_PATH}/protocol/service/$SRV/*.proto
	# generating grpc service stub
	protoc -I/protobuf \
		-I${ROOT_PATH} \
		--go_out=plugins=grpc:${GO_OUTPUT} \
		${FILE_PATH}
	# generating gateway map
	protoc -I/protobuf \
		-I${ROOT_PATH} \
        --grpc-gateway_out=logtostderr=true:${ROOT_PATH} \
		${FILE_PATH}
	# generating swagger documentation
	protoc -I/protobuf \
		-I${ROOT_PATH} \
        --swagger_out=logtostderr=true:${ROOT_PATH} \
		${FILE_PATH}
done
