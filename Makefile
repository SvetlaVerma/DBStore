CURDIR:=$(shell pwd)

test:
	cd ${CURDIR};go test -v ./...

clean:
	cd ${CURDIR}
	rm -f ${SVCAPPNAME}

protoc:
	cd ${CURDIR}/MicroServices;
	protoc -I grpc/ grpc/db.proto --go_out=plugins=grpc:grpc

run:
	cd ${CURDIR}/MicroServices/DBService;go build; ./DBService &
	cd ${CURDIR}/MicroServices/CSVParser;go run main.go;cd -;
