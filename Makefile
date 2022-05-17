.PHONY: setup
setup:
	go install github.com/cosmtrek/air@latest

.PHONY: env
env:
	cp develop.env .envrc
	direnv allow .

.PHONY: spanner
spanner: 
	docker-compose exec spanner-cli spanner-cli -p ${SPANNER_PROJECT_ID} -i ${SPANNER_INSTANCE_ID} -d ${SPANNER_DATABASE_ID}