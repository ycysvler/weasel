assets:
	@echo "--> Running assets"
	@go-bindata -ignore \.go -pkg schema -o ./schema/bindata.go ./schema/...
	@cd $(CURDIR)
	@go fmt $$(go list ./... | grep -v /schema)
	@echo "$(G)[OK]$(C)"


