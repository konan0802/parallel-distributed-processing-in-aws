FROM golang:alpine AS build-stage
ADD . /work
WORKDIR /work
RUN GOOS=linux GOARCH=amd64 go build -o main

FROM alpine:latest
COPY --from=build-stage /work/main /usr/local/bin/main
ENTRYPOINT ["/usr/local/bin/main"]