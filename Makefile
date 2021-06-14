
# This makefile defines the following targets
#
#   - all (default) - formats the code and downloads vendor libs
#   - fmt - formats the code
#   - vendor - download all third party libraries and puts them inside vendor directory
#   - clean-vendor - removes third party libraries from vendor directory

.PHONY: all				##formats the code and downloads vendor libs
all: fmt vendor

.PHONY: fmt				##formats the code
fmt:
	@go fmt ./...

.PHONY: vendor			##download all third party libraries and puts them inside vendor directory
vendor:
	@go mod vendor

.PHONY: clean-vendor			##removes third party libraries from vendor directory
clean-vendor:
	-@rm -rf vendor

.PHONY: help				##show this help message
help:
	@echo "usage: make [target]\n"; echo "options:"; \fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//' | sed 's/.PHONY:*//' | sed -e 's/^/  /'; echo "";