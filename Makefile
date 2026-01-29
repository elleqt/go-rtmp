.PHONY: vet
vet:
	go vet $$(go list ./... | grep -v /vendor/)

.PHONY: test
test:
	go test -cover -coverprofile=coverage.txt -covermode=atomic -v -race -timeout 10s ./...

.PHONY: bench
bench:
	go test -bench . -benchmem -gcflags="-m -m -l" ./...
