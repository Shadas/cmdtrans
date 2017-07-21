build : main.go
	mkdir -p bin
	go build -i -o bin/fy cmdtrans
	GOOS=darwin  GOARCH=amd64 go build -i -o bin/fy-Darwin-x86_64 cmdtrans
	GOOS=linux   GOARCH=amd64 go build -i -o bin/fy-Linux-x86_64 cmdtrans 
	GOOS=windows GOARCH=amd64 go build -i -o bin/fy.exe cmdtrans

clean :
	rm -rf bin