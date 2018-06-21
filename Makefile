schema: ;
	go-bindata -ignore=\.go -pkg=schema -o=schema/bindata.go ./...

.PHONY: schema
