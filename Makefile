all:
	go build cmd/lambda/main.go 
	cd tests && make