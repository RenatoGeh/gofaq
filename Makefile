.PHONY: test
test:
	go build -o test/gofaq; cp example.faq test/; cp imgs test/ -r; cd test; ./gofaq -url=${PWD}/test; cd ..

build:
	go build -o test/gofaq

.PHONY: clean
clean:
	rm test/* -r
