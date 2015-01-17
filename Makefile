PREFIX=/usr/local
RUNTIME_GOPATH=$(GOPATH):`pwd`
VERSION=`git tag | tail -n 1`
GOOS=`go env GOOS`
GOARCH=`go env GOARCH`

banner:	main.go src/banner/banner.go
	GOPATH=$(RUNTIME_GOPATH) go build -o banner main.go

install: banner
	install -m 755 banner $(PREFIX)/bin/

clean:
	rm -f banner *.tar.gz

test:
	GOPATH=$(RUNTIME_GOPATH) go test src/**/*_test.go

package: clean banner
	tar zcf sysvbanner-go-$(VERSION)-${GOOS}-$(GOARCH).tar.gz ./banner
