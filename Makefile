
test:
	cd ${CURDIR};go test -v ./...

clean:
	cd ${CURDIR}
	rm -f ${SVCAPPNAME}

proto:
    cd ${CURDIR}/grpc;
    protoc -I dbrpc/ dbrpc/db.proto --go_out=plugins=grpc:dbrpc

run:
	cd ${CURDIR}/DBService;
	go run main.go
	cd -
	cd ${CURDIR}/CSVParser;
	go run main.go
	cd -