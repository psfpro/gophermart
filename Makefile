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
	gophermarttest \
            -test.v -test.run=^TestGophermart$$ \
            -gophermart-binary-path=cmd/gophermart/gophermart \
            -gophermart-host=localhost \
            -gophermart-port=8080 \
            -gophermart-database-uri="postgresql://app:pass@localhost:5432/app?sslmode=disable" \
            -accrual-binary-path=cmd/accrual/accrual_linux_amd64 \
            -accrual-host=localhost \
            -accrual-port=8081 \
            -accrual-database-uri="postgresql://app:pass@localhost:5432/app?sslmode=disable"

