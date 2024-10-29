GOPROXY=GOPROXY=https://internal-gomod.milvzn.com GONOSUMDB='*' GODEBUG=madvdontneed=1


## mod_tidy: will run go mod tidy
mod_tidy:
	@$(GOPROXY) go mod tidy
