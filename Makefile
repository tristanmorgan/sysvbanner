PREFIX=/usr/local

banner:	main.go src/banner/banner.go
	GOPATH=$(GOPATH):`pwd` go build -o banner main.go

install: banner
	install -m 755 banner $(PREFIX)/bin/

clean:
	rm -f banner
