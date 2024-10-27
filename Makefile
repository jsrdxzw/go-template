BIN_SRC=cmd/main.go
DIST_DIR=.dist
LOG_FILE=*.log*

##################################
#######       Setup       ########
##################################
.PHONY: ensure tidy

tidy:
	@go mod tidy

ensure: tidy
	@go mod download

##################################
#######        Tool       ########
##################################
.PHONY: fmt lint build clean api

API_PATH = "./service/api/example.api"
API_DIR = "./service/api"
API_TEMPLATE = "./.goctl"
SWAGGER_GEN_DIR = "./gen/swagger"
ifeq ($(shell uname),Darwin)
    CMD = gsed
else
    CMD = sed
endif

api:
	go install github.com/zeromicro/go-zero/tools/goctl@latest
	goctl api format --dir ${API_DIR} --declare
	goctl api go --api ${API_PATH} --dir ${API_DIR}
	find * -type f -exec ${CMD} -i 's/^\s\+$$//g' {} +

##################################
#######     Migrate       ########
##################################
.PHONY: migrations migrate sqlmigrate

SCHEMA_DIR=./spec/schema

SCHEMA_DIR=./spec/schema

migrations:
	go install entgo.io/ent/cmd/ent@v0.14.1
	ent generate --feature sql/upsert --target gen/entschema ${SCHEMA_DIR}

migrate:
	@go run ${BIN_SRC} migrate

sqlmigrate:
	@go run ${BIN_SRC} migrate --fake
