assets:
	@echo "--> Running assets"
	@go-bindata -ignore \.go -pkg schema -o ./schema/bindata.go ./schema/...
	@cd $(CURDIR)
	@go fmt $$(go list ./... | grep -v /schema)
	@echo "$(G)[OK]$(C)"

tools:
	@echo "--> Running tools"
	@go get $(GOTOOLS)
	@echo "$(G)[OK]$(C)"

dev: assets
	@go run main.go