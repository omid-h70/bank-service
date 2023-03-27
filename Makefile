SHELL=cmd.exe
APP_BINARY=bankService.exe

##
## Attention
## All MakeFile Command Lines Must Contain Tab Not Spaces
##
##

##
## CROSS Compiling Golang
## By Setting These Options we're cross Compiling Go From Windows To Linux
## GOOS=linux && set GOARCH=amd64
##

## up: starts all containers in the background without forcing build
up:
	@echo "Starting Docker images..."
	docker-compose up -d
	@echo "Docker images started!"

### up_build: stops docker-compose (if running), builds all projects and starts docker compose
#up_build: build_broker
#    @echo Stopping docker images (if running...)
#    docker-compose down
#    @echo Building (when required) and starting docker images...
#    docker-compose up --build -d
#    @echo Docker images built and started!
#
### down: stop docker compose
down:
	@echo "Stopping docker compose..."
	docker-compose down
	@echo "Done!"
#
### build_app: builds the broker binary as a linux executable
## Append && to GOOS=linux File !!!!!!!!!!! with no spaces !!!!!!!!!!!!
## && means execute second command if first was successful
build_app:
	@echo "Building Bank Service..."
	chdir ..\bank-service && set GOOS=linux&& set GOARCH=amd64&& set CGO_ENABLED=0 && go build -o ${APP_BINARY} .
	@echo "Done!"

### start: starts the front end
#start: build_front
#    @echo Starting front end
#    chdir ..\front-end && start /B ${FRONT_END_BINARY} &
#
### stop: stop the front end
#stop:
#    @echo Stopping front end...
#    @taskkill /IM "${FRONT_END_BINARY}" /F
#    @echo "Stopped front end!"
