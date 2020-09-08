BINARY=meroedu
DB_CONFIG_FILE ?= ./migrator/dbconf.yml
DB_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(DB_CONFIG_FILE))
##############################################################################
# Development/Staging
##############################################################################
config:
	cp config.example.yml config.yml && cd migrator && cp dbconf.example.yml dbconf.yml
wait-for-db:
	./scripts/wait-for-db.sh
	sleep 5
init-db:
	docker-compose up -d db

prepare: config init-db wait-for-db migrate

run:
	docker-compose up --build -d
stop:
	docker-compose down
reset: migrate-down stop
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

test: 
	go test -v -cover -covermode=atomic ./...

unittest:
	go test -short  ./...

test-coverage:
	mkdir -p ./out
	go test -coverprofile=./out/coverage.out ./...
	go tool cover -func=./out/coverage.out

sonar: test
	sonar-scanner -Dsonar.projectVersion="$(version)"

start-sonar:
	docker run --name sonarqube -p 9000:9000 sonarqube
#############################################################################
# Migration
#############################################################################
migrate-build:
	cd migrator/ && docker build -t migrator .

migrate: migrate-build
	docker run --network host migrator -path=/migrations/ -database "$(DB_DSN)" up

migrate-down:
	docker run --network host migrator -path=/migrations/ -database "$(DB_DSN)" down -all
	
#############################################################################
# Utility
#############################################################################
build:
	go build -o ${BINARY}
build-app: clean-app
	go build -o ${BINARY}

run-app: build-app
	./${BINARY}
clean-app:
	$(eval VALUE=$(shell sh -c "lsof -i:9090 -t"))
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	$(shell sh -c "if [ \"${VALUE}\" != \"\" ]  ; then kill ${VALUE} ; fi")
docker:
	docker build -t course_api .

swagger:
	go get github.com/swaggo/swag/cmd/swag
	$$(go env GOPATH)/bin/swag init -g meroedu.go --output ./api_docs
mock:
	cd internal/domain && mockery --all --keeptree
db-up:
	docker-compose up -d mysql
.PHONY: clean install unittest build docker run stop vendor lint-prepare lint
