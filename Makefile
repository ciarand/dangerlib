compile:
	go build server/receiver.go
	go build client/sender.go
	rm /tmp/build.sock
