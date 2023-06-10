GQLGEN := $(shell grep github.com/99designs/gqlgen go.mod | tr ' ' '@')

generate:
	go run $(GQLGEN) generate
.PHONY: generate