build: build-server build-client

build-server:
        docker build -t server:v1 -f build/package/server/Dockerfile .

build-client:
        docker build -t client:v1 -f build/package/client/Dockerfile .

run:
        docker-compose up --force-recreate