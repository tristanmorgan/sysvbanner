PREFIX=/usr/local
RUNTIME_GOPATH=$(GOPATH):`pwd`

banner:	main.go src/banner/banner.go
	GOPATH=$(RUNTIME_GOPATH) go build -o banner main.go

install: banner
	install -m 755 banner $(PREFIX)/bin/

clean:
	rm -f banner

test:
	GOPATH=$(RUNTIME_GOPATH) go test src/**/*_test.go
