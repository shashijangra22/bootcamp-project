# End point for DB
EP = --endpoint http://localhost:8000

# Table names
TABLE1 = assets/db_schema/customers.json
TABLE2 = assets/db_schema/orders.json
TABLE3 = assets/db_schema/restaurants.json

# proto files to generate stub code
PROTO1 = pkg/protos/customer.proto
PROTO2 = pkg/protos/order.proto
PROTO3 = pkg/protos/restaurant.proto

# rule to create all tables above
tables:
	aws dynamodb create-table --cli-input-json file://$(TABLE1) $(EP)
	aws dynamodb create-table --cli-input-json file://$(TABLE2) $(EP)
	aws dynamodb create-table --cli-input-json file://$(TABLE3) $(EP)

# rule to generate stub code from proto files
protos:
	protoc $(PROTO1) --go_out=plugins=grpc:.
	protoc $(PROTO2) --go_out=plugins=grpc:.
	# protoc $(PROTO3) --go_out=plugins=grpc:.

# rule to populate DB with sample data
populate:
	go run cmd/DB/main.go assets/sample_data

# rule to run gin-API-Client
client:
	go run cmd/GIN-CLIENT/main.go

# rule to run gRPC-Server
server:
	go run cmd/SERVER/main.go

# rule to list all tables in DB
list-tables:
	aws dynamodb list-tables $(EP)

# rule to describe a table in DB
show-table:
	aws dynamodb describe-table --table-name $(T) $(EP)

# rule to delete a given table
delete-table:
	aws dynamodb delete-table --table-name $(T) $(EP)