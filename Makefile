
test:
	go test -v -coverprofile=coverage.txt -covermode=atomic .

clean:
	rm -f coverage.txt
