CURDIR=$(shell pwd)
BINDIR=${CURDIR}/bin
GOVER=$(shell go version | perl -nle '/(go\d\S+)/; print $$1;')
SMARTIMPORTS=${BINDIR}/smartimports_${GOVER}
LINTVER=v1.51.1
LINTBIN=${BINDIR}/lint_${GOVER}_${LINTVER}
GOOSEBIN=${BINDIR}/goose
PACKAGE=github.com/Arkosh744/banners

LOCAL_MIGRATION_DIR=./migrations
LOCAL_MIGRATION_DSN="host=localhost port=6662 dbname=banners-db user=banners-user password=banners-pass sslmode=disable"

new-migration:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} create new-migration sql

local-migration-status:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} status -v

local-migration-up:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} up -v

local-migration-down:
	${GOOSEBIN} -dir ${LOCAL_MIGRATION_DIR} postgres ${LOCAL_MIGRATION_DSN} down -v

build: bindir
	go build -o ${BINDIR}/app ${PACKAGE}

test-integration:
	go test -tags=integration -cover ./...

test:
	go test -cover ./...

test-all: test-integration test

run:
	go run ${PACKAGE}

lint: install-lint
	${LINTBIN} run

bindir:
	mkdir -p ${BINDIR}

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
