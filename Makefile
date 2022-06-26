include .env
export $(shell sed 's/=.*//' .env)

compose-up:
	docker compose up -d

run:
	go run .

.PHONY: test
test:
	go test -v ./...

ec:
	echo "blas"