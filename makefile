build: 
	go build -o bin/GO-REST

run:	build	
	./bin/GO-REST

test:
	@go test -v ./..