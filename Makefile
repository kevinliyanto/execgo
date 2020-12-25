build:
	go build -mod vendor

all:
	build

clean:
	rm execgo