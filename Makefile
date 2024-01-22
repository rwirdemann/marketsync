build-web:
	packr
	go build -o bin/ms context/web/main.go
	packr clean

deploy: build-web
	cp bin/ms /Users/ralf/Documents/work/electron/marketsync-el
	cp bestand.xlsx /Users/ralf/Documents/work/electron/marketsync-el
	cp marketsync.yaml /Users/ralf/Documents/work/electron/marketsync-el

clean:
	rm -rf ./bin

.PHONY: build-web clean