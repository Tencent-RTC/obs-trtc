.PHONY: default clean

default: ./main

clean:
	rm -f main

./main: *.go Makefile go.mod
	go build .

