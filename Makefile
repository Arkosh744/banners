CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
GOOSEBIN=${BINDIR}/goose
PACKAGE=github.com/Arkosh744/banners/cmd/server

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=5439 dbname=banners-db user=banners-user password=banners-pass sslmode=disable"

bindir:
	mkdir -p ${BINDIR}

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

run: build
	sudo docker compose up --force-recreate --build -d
	make local-migration-up

new-migration:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} create new-migration sql

local-migration-status:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

install-go-deps:
	mkdir -p bin
	GOBIN=$(BINDIR) go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.1
	GOBIN=$(BINDIR) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

generate:
	mkdir -p pkg/banners_v1
	protoc --proto_path api/banners_v1 --proto_path vendor.protogen \
	--go_out=pkg/banners_v1 --go_opt=paths=source_relative \
	--plugin=protoc-gen-go=bin/protoc-gen-go \
	--go-grpc_out=pkg/banners_v1 --go-grpc_opt=paths=source_relative \
	--plugin=protoc-gen-go-grpc=bin/protoc-gen-go-grpc \
	api/banners_v1/banners.proto

test-integration:
	go test -tags=integration -cover ./...

test:
	go test -cover ./...

test-all: test-integration test

lint: install-lint
	gofumpt -w -extra ./
	${LINTBIN} run --fix

format: install-smartimports
	${SMARTIMPORTS} -exclude internal/mocks

install-all: install-lint install-goose install-smartimports

install-lint: bindir
	test -f ${LINTBIN} || \
		(GOBIN=${BINDIR} go install github.com/golangci/golangci-lint/cmd/golangci-lint@${LINTVER} && \
		mv ${BINDIR}/golangci-lint ${LINTBIN})

install-goose: bindir
	test -f ${GOOSEBIN} || GOBIN=${BINDIR} go install github.com/pressly/goose/v3/cmd/goose@latest

install-smartimports: bindir
	test -f ${SMARTIMPORTS} || \
		(GOBIN=${BINDIR} go install github.com/pav5000/smartimports/cmd/smartimports@latest && \
		mv ${BINDIR}/smartimports ${SMARTIMPORTS})

vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi