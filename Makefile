init:
	cd go-grpc-auth-svc && go get
	cd API-Geteway && go get
	cd go-grpc-product-svc && go get
	cd go-grpc-order-svc && go get
proto:
	cd go-grpc-auth-svc && make proto 
	cd API-Geteway && make proto
	cd go-grpc-product-svc && make proto
	cd go-grpc-order-svc && make proto
test: 
	cd go-grpc-auth-svc && make cover && make coverClean
postgresCreateDB:
	docker-compose -f docker-compose.dev.yml up -d postgres
	docker cp ./init.sql postgres-container:/init.sql
	docker exec -it postgres-container psql -U postgres -W -a -v ON_ERROR_STOP=1 -f ./init.sql
dockerRun:
	docker-compose -f docker-compose.dev.yml up -d
