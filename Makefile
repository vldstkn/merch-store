api:
	@air -c air/.api.toml
acc:
	@air -c air/.account.toml
prod:
	@air -c air/.products.toml
tran:
	@air -c air/.transfers.toml

pb:
ifdef s
	$(MAKE) gen SERVICE=$(s)
else
	$(MAKE) gen SERVICE=account
	$(MAKE) gen SERVICE=products
	$(MAKE) gen SERVICE=transfers
endif
gen:
	@protoc \
		--proto_path=proto "./proto/$(SERVICE).proto" \
		--go_out=pkg/pb \
		--go_opt=paths=source_relative \
		--go-grpc_out=pkg/pb \
		--go-grpc_opt=paths=source_relative