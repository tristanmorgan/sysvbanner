PREFIX=/usr/local

banner:	banner.go
	go build banner.go

install:
	install -m 755 banner $(PREFIX)/bin/

clean:
	rm -f banner
