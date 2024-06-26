kitex:
	@rm -rf kitex_gen/
	@kitex -module github.com/zlllgp/vegas -service vegas ./idl/vegas.proto
	@rm build.sh
	@rm -rf script