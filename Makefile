
test:
	-rm _test_tmp/*.db
	godep go test -v .

build:
	godep go build .
