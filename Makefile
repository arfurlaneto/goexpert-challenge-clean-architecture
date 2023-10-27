.PHONY: update-wire
update-wire:
	@echo Updating wire files...
	@cd cmd/ordersystem;	wire

.PHONY: update-grpc
update-grpc:
	@echo Updating gRPC files...
	@protoc --go_out=. --go-grpc_out=. internal/infra/grpc/protofiles/order.proto

.PHONY: update-graphql
update-graphql:
	@echo Updating GraphQL files...
	go run github.com/99designs/gqlgen generate

.PHONY: run-migrations
run-migrations:
	@echo Running database migrations
	@docker exec mysql mysql -uroot -proot orders -e "CREATE TABLE IF NOT EXISTS orders (id varchar(255) NOT NULL, price float NOT NULL, tax float NOT NULL, final_price float NOT NULL, PRIMARY KEY (id));"

.PHONY: run
run:
	@echo Starting application...
	@cd cmd/ordersystem; go run main.go wire_gen.go
