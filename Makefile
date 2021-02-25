test:
	go run ./cmd/forkinternal cmd/go/internal/modload cmd/go/internal/modfetch
	tree ./forked/

install:
	go install ./cmd/forkinternal

dep:
	go get -u ./...