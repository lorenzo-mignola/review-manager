run-server:
	go run ./server/main.go

run-client:
	cd client && yarn dev

run-all: run-server run-client