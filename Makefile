GOCMD			=	go
GOFMT			=	$(GOCMD) fmt
GOFUMPT		=	gofumpt
GOVET			=	$(GOCMD) vet
GOBUILD		=	$(GOCMD) build
GORETURNS	= goreturns
LINTER		=	golangci-lint
APP				=	authenticator-backend
PLATFORM	=	linux/arm64
ENV				= local
PORT			= 8081
CONTAINER	= ko.local/$(APP)
PGURL			=	postgres://dhuser:passw0rd@db:5432/dhlocal?sslmode=disable
DOCKER_NW	=	docker.internal
MOCK_SRC_REPOSITORY = $(wildcard domain/repository/*.go)
MOCK_SRC_USECASE = $(wildcard usecase/*usecase.go)
MOCK_SRC_HANDLER = $(wildcard presentation/http/echo/handler/*.go)
MOCK_FILES = $(wildcard test/mock/*.go)

export KO_DEFAULTBASEIMAGE	:=	debian:bullseye-slim

.PHONY: test

all:
	make goreturns
	make genmock
	make vet
	make lint-fix
	make test
	make swaggo
	make build-go
	make build-with-docker
	make scan-image

validate:
	$(GOFMT) -w -s ./...

goreturns:
	$(GORETURNS) -w .

lint:
	$(LINTER) run

lint-fix:
	$(LINTER) run --fix

vet:
	$(GOVET) ./...

swaggo:
	swag init

build:
	ko publish --local --platform=$(PLATFORM) --base-import-paths .

build-with-docker:
	docker build -t $(APP) .

build-go:
	$(GOBUILD) main.go

genmock: $(MOCK_SRC_REPOSITORY)
	rm $(MOCK_FILES)
	go generate $(MOCK_SRC_REPOSITORY)
	go generate $(MOCK_SRC_USECASE)
	go generate $(MOCK_SRC_HANDLER)

test:
	go test -v -cover -covermode=atomic ./...

test-coverage:
	go test -v -cover -coverprofile=cover.out -covermode=atomic ./presentation/http/echo/handler/... ./usecase/... ./infrastructure/persistence/datastore/... ./infrastructure/gocloak/...
	go tool cover -html=cover.out -o cover.html

integration-test:
	./test/integration_test/run.sh $(ENV) $(TARGET)

run:
	docker run -v $(PWD)/config/:/app/config/ -td -i --network docker.internal --env-file config/$(ENV).env -p $(PORT):$(PORT) --name $(APP) $(APP)

run-with-ko:
	docker run -td -i --network docker.internal --env-file config/$(ENV).env -p $(PORT):$(PORT) --name $(APP) $(CONTAINER)

run-go:
	set -a && . config/local-developer.env && set +a && go run main.go

run-air:
	set -a && . config/local-developer.env && set +a && air

run-memory:
	export DATASTORE=memory && set -a && . config/$(ENV).env && set +a && air

tbls:
	tbls doc --force "$(PGURL)"

scan-image:
	docker run -v /var/run/docker.sock:/var/run/docker.sock --rm aquasec/trivy image --severity HIGH,CRITICAL $(APP)

clean:
	docker stop $(APP); docker rm $(APP)
	docker container prune --force

api-scan:
	./scripts/api-scan.sh

db-dev:
	gcloud compute ssh data-spaces-bastion-dev \
	--project data-spaces-dev-j8wlpu7j \
	--zone us-west1-a --tunnel-through-iap -- -N -L 1234:localhost:1234

db-sbx:
	gcloud compute ssh data-spaces-bastion-sbx \
	--project data-spaces-sbx-c70v3nhf \
	--zone us-west1-a --tunnel-through-iap -- -N -L 1234:localhost:1234

db-sbx2:
	gcloud compute ssh data-spaces-bastion-sbx2 \
	--project data-spaces-sbx2-6obg6a6w \
	--zone us-west1-a --tunnel-through-iap -- -N -L 1234:localhost:1234

db-qa:
	gcloud compute ssh data-spaces-bastion-qa \
	--project data-spaces-qa-piqrxp1z \
	--zone us-west1-a --tunnel-through-iap -- -N -L 1234:localhost:1234

db-stg:
	gcloud compute ssh data-spaces-bastion-stg \
	--project data-spaces-stg-pa8qxolm \
	--zone asia-northeast1-b --tunnel-through-iap -- -N -L 1234:localhost:1234

keycloak-add-local:
	go run cmd/add_local_user/main.go

keycloak-add-tutorial:
	go run cmd/add_tutorial_user/main.go