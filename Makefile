build-web:
	packr
	go build -o bin/ms context/web/main.go
	packr clean

clean:
	rm -rf ./bin

.PHONY: build-web clean