.PHONY: generate

generate:
	oapi-codegen --config=./config.yml ./openapi.yml