client-build:
	go build -o ./bin/client ./cmd/client/main.go

client-run: client-build
	./bin/client

server-build:
	go build -o ./bin/server ./cmd/server/main.go

server-run: server-build
	./bin/server

up:
	docker compose up -d --build

down:
	docker compose down -v --rmi all

clean:
	rm -rf bin

test:
	go test -v ./... -count=1
