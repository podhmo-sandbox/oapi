generate:
	go run ./tools/gen.go > openapi.json
.PHONY: generate

validate:
	go run github.com/getkin/kin-openapi/cmd/validate@latest ./openapi.json
.PHONY: validate
