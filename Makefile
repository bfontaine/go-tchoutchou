SRCS=cli.go $(wildcard tchou/*.go)

bin/tchou: $(SRCS)
	go build -o $@ .

check:
	go test -v ./...
