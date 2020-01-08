all: bin
	go build -o bin -v ./cmd/...

bin:
	mkdir bin

install:
	go install -v ./cmd/...

clean:
	rm ./bin/*
