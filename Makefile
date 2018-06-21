schema: ;
	go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go ./...
coverage:
	go test -v -coverprofile=c.out ./...;
	go tool cover -html=c.out;
	rm c.out

.PHONY: schema coverage
