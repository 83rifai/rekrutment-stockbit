# PROJECT_ROOT=gitlab.com/transdigital/engineering/grpc-pim
PROTO_FILES=$(shell find . -path '*.proto' | grep proto)

proto-gen:
	@$(foreach file,$(PROTO_FILES),protoc $(file) -I. --go_out=plugins=grpc:.;)

print:
	@$(foreach file,$(PROTO_FILES),protoc $(file) -I. --go_out=plugins=grpc:. && echo $(file);)