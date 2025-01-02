PACKAGE = $(shell head -1 go.mod | awk '{print $22}')
build: 	generate
	go build -o ${BIN} .

generate:
	protoc -I${PROTO_DIR} --go_opt=module=${PACKAGE} --go_out=. ${PROTO_DIR}/*.proto

bump: generate
	go get -u ./...

clean:
	rm ${PROTO_DIR}/*.pb.go
	rm ${BIN}