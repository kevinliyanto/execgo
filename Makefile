build:
	go build -mod vendor -o output

all:
	build

clean:
	rm output