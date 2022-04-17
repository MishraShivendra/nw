FROM golang:1.17 as build

WORKDIR /src
COPY . .

RUN cd /src/server;go mod tidy; go build server.go
RUN cd /src/client;go mod tidy; go build client.go

FROM ubuntu:latest as runer
COPY run.sh .
COPY yaml/users.yaml ./users.yaml
COPY --from=build /src/client/client ./client
COPY --from=build /src/server/server ./server
CMD ["/bin/bash", "run.sh"]
