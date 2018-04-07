#/bin/bash

~/jungleCode/bin/protoc --go_out=plugins=grpc:. addrbookstore.proto
