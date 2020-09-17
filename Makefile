.PHONY: build
build: test imei

.PHONY: test
test: *.go
	go test ./...

.PHONY: clean
clean:
	rm imei

imei: *.go cli/imei/*.go
	go build ./cli/imei/...
