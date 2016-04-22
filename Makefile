default: build

pwd := $(shell pwd)

build:
	docker run \
		-e CGO_ENABLED=0 \
		-e GOOS=linux \
		-v $(GOPATH)/src/:/go/src/ \
		-v $(pwd):/go/src/github.com/serdmanczyk/freyr/ \
		-w /go/src/github.com/serdmanczyk/freyr/ \
		golang go build -ldflags "-s" -a -installsuffix cgo -o freyr

test:
	docker-compose -f docker-compose.test.yml -p ci up --force-recreate

sql:
	docker-compose -f docker-compose.sql.yml -p sql down
	docker-compose -f docker-compose.sql.yml -p sql up --force-recreate

rundev:
	docker-compose -p dev down
	docker-compose -f docker-compose.debug.yml -p dev build
	docker-compose -f docker-compose.debug.yml -p dev up --force-recreate

runstatic:
	docker-compose build
	docker-compose up
