gen:
	protoc -I ./proto/deposit \
       --go_out ./proto/deposit --go_opt paths=source_relative \
       --go-grpc_out ./proto/deposit --go-grpc_opt paths=source_relative \
       ./proto/deposit/deposit.proto