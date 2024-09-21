test *ARGS:
    go test ./... {{ARGS}}

bench *ARGS:
    go test -bench . -benchmem ./... {{ARGS}}
