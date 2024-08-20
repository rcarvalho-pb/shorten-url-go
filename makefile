BINARY_NAME=myapp
DSN=db/db.db
PORT=8080
REDIS="127.0.0.1:6379"

## build: Build binary
build:
	@echo "Building..."
	go build -ldflags="-s -w" -o ${BINARY_NAME} ./cmd
	@echo "Built!"

## run: builds and runs the application
run: build
	@echo "Starting..."
	@env DSN=${DSN} PORT=${PORT} ./${BINARY_NAME} &
	@echo "Started!"

## clean: runs go clean and deletes binaries
clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

## start: an alias to run
start: run

## stop: stops the running application
stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

## restart: stops and starts the application
restart: stop start

## test: runs all tests
test:
	go test -v ./...
