BINARY=engine
##############################################################################
# Staging
##############################################################################

run:
	docker-compose -f docker-compose.yaml up --build -d
stop:
	docker-compose -f docker-compose.yaml down
##############################################################################
# Development
##############################################################################

run-dev:
	docker-compose -f docker-compose.dev.yaml up --build
stop-dev:
	docker-compose -f docker-compose.dev.yaml down

##############################################################################
# Lint
###############################################################################
lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

##############################################################################
# Test
###############################################################################

test-with-richgo: 
	richgo test -v -cover -covermode=atomic ./...

test: 
	go test -v -cover -covermode=atomic ./...

unittest:
	go test -short  ./...

#############################################################################
# Migration
#############################################################################
migrate-build:
	cd migrator/ && docker build -t migrator .

migrate: migrate-build
	docker run --network host migrator -path=/migrations/ -database "mysql://user:password@tcp(127.0.0.1:3306)/course_api" up 1

migrate-down:
	docker run --network host migrator -path=/migrations/ -database "mysql://user:password@tcp(127.0.0.1:3306)/course_api" down 1
	
#############################################################################
# Utility
#############################################################################
engine:
	go build -o ${BINARY} app/*.go

run-engine:
	./engine
clean:
	$(eval VALUE=$(shell sh -c "lsof -i:9090 -t"))
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	$(shell sh -c "if [ \"${VALUE}\" != \"\" ]  ; then kill ${VALUE} ; fi")
docker:
	docker build -t course_api .
db-up:
	docker-compose up -d mysql
.PHONY: clean install unittest build docker run stop vendor lint-prepare lint