GORESTAPI_FOLDER = ./gorestapi
CONFIG_FOLDER = ./config

build: $(GORESTAPI_FOLDER)
	CONFIG_FOLDER=$(CONFIG_FOLDER) go build -a -o bin/gorestapi run/server.go

run: $(GORESTAPI_FOLDER)
	CONFIG_FOLDER=$(CONFIG_FOLDER) go run run/server.go

test: $(GORESTAPI_FOLDER)
	CONFIG_FOLDER=../$(CONFIG_FOLDER) go test -v $(GORESTAPI_FOLDER)

