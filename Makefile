kitex:
	@rm -rf kitex_gen/
	@kitex -module github.com/zlllgp/vegas -service vegas ./idl/vegas.proto
	@rm build.sh
	@rm kitex_info.yaml
	@rm -rf script