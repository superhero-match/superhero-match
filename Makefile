prepare:
	go mod download

run:
	go build -o bin/main cmd/api/main.go
	./bin/main

build:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o bin/main cmd/api/main.go
	chmod +x bin/main

dkb:
	docker build -t superhero-match .

dkr:
	docker run -p "4200:4200" -p "8180:8180" superhero-match

launch: dkb dkr

api-log:
	docker logs superhero-match -f

es-log:
	docker logs es -f

rmc:
	docker rm -f $$(docker ps -a -q)

rmi:
	docker rmi -f $$(docker images -a -q)

clear: rmc rmi

api-ssh:
	docker exec -it superhero-match /bin/bash

es-ssh:
	docker exec -it es /bin/bash

PHONY: prepare build dkb dkr launch api-log es-log api-ssh es-ssh rmc rmi clear