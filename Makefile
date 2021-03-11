-include config
PROTO_FILES=$(shell find . -path '*.proto' | grep proto)

run:
	@go run main.go

proto-gen:
	@$(foreach file,$(PROTO_FILES),protoc $(file) -I. --go_out=plugins=grpc:.;)

print:
	@$(foreach file,$(PROTO_FILES),protoc $(file) -I. --go_out=plugins=grpc:. && echo $(file);)