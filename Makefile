.PHONY: setup
setup:
	bash ./script/setup.sh
.PHONY: env
env:
	cp develop.env .envrc
	direnv allow .

.PHONY: spanner
spanner: 
	docker-compose exec spanner-cli spanner-cli -p ${SPANNER_PROJECT_ID} -i ${SPANNER_INSTANCE_ID} -d ${SPANNER_DATABASE_ID}

.PHONY: createdb
createdb:
	wrench create --directory ./ddl
	wrench migrate up --directory .

.PHONY: resetdb
resetdb:
	wrench reset --directory ./ddl
	wrench migrate up --directory ./ddl

.PHONY: lint 
lint:
	# golangci-lint run ./...
	# go vet -vettool=$(shell which zagane) github.com/uyuhub/cloud-spanner/infra/spanner


