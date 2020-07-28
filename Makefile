BINARY=engine
test-with-richgo: 
	richgo test -v -cover -covermode=atomic ./...

test: 
	go test -v -cover -covermode=atomic ./...

engine:
	go build -o ${BINARY} app/*.go

run-engine:
	./engine

unittest:
	go test -short  ./...

clean:
	$(eval VALUE=$(shell sh -c "lsof -i:9090 -t"))
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi
	$(shell sh -c "if [ \"${VALUE}\" != \"\" ]  ; then kill ${VALUE} ; fi")

docker:
	docker build -t course_api .
db-up:
	docker-compose up -d mysql
run:
	docker-compose -f docker-compose.yaml up --build -d
stop:
	docker-compose -f docker-compose.yaml down

run-dev:
	docker-compose -f docker-compose.dev.yaml up --build
stop-dev:
	docker-compose -f docker-compose.dev.yaml down

lint-prepare:
	@echo "Installing golangci-lint" 
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest

lint:
	./bin/golangci-lint run ./...

.PHONY: clean install unittest build docker run stop vendor lint-prepare lint