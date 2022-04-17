FROM golang:1.17 as build

WORKDIR /src
COPY . /src
RUN go mod tidy
RUN cd /src/server; go build server.go
RUN cd /src/client; go build client.go

FROM ubuntu:latest as runer
COPY --from=build /src/client/client ./client
COPY --from=build /src/server/server ./server
CMD ["/bin/bash", "run.sh"]
