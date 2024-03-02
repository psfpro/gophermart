up:
	docker compose up -d

down:
	docker compose down

test: clear vet build-gophermart gophermart-test

clear:
	clear

vet:
	go vet ./...

build-gophermart:
	cd cmd/gophermart && go build -buildvcs=false -o gophermart

gophermart-test:
	metricstest -test.v -test.run=^TestIteration14$$ \
                -source-path=. \
                -agent-binary-path=cmd/agent/agent \
                -binary-path=cmd/server/server \
                -database-dsn='postgres://app:pass@localhost:5432/app?sslmode=disable' \
                -key=123 \
                -file-storage-path=tmp \
                -server-port=8888
