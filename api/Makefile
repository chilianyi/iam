# Copyright 2018 The OpenPitrix Authors. All rights reserved.
# Use of this source code is governed by a Apache license
# that can be found in the LICENSE file.

PWD:=$(shell pwd)
GOPATH:=$(shell go env GOPATH)

PBFILES=$(sort $(wildcard *.proto))

PROTOC_FLAGS:=-I. -Ithird_party/googleapis -Ithird_party

generate: Makefile
	@rm -rf ../pkg/pb/*
	@mkdir -p ../pkg/pb

	# grpc service
	protoc $(PROTOC_FLAGS) --go_out=plugins=grpc:$(GOPATH)/src ${PBFILES}

	# govalidators
	#protoc $(PROTOC_FLAGS) --govalidators_out=$(GOPATH)/src ${PBFILES}

	# gateway
	protoc $(PROTOC_FLAGS) --grpc-gateway_out=logtostderr=true,allow_delete_body=true:../pkg/pb ${PBFILES}

	# swagger
	protoc $(PROTOC_FLAGS) --swagger_out=logtostderr=true,allow_delete_body=true:../pkg/service/spec ${PBFILES}
	make -C ../pkg/service/spec

clean:
	rm -rf $(GOPATH)/src/openpitrix.io/iam/pkg/pb