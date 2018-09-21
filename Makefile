CURDIR:=$(shell pwd)

test:
	cd ${CURDIR};go test -v ./...

clean:
	cd ${CURDIR}
	rm -f ${SVCAPPNAME}

proto: cd ${CURDIR}/grpc;
    protoc -I dbrpc/ dbrpc/db.proto --go_out=plugins=grpc:dbrpc

run:
	cd ${CURDIR}/MicroServices/DBService;go build; ./DBService &
	cd ${CURDIR}/MicroServices/CSVParser;go run main.go;cd -;
