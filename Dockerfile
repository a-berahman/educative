FROM golang:1.15 AS build

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN GOPATH=/usr/go CGO_ENABLED=0 go build -o educative .

FROM alpine:3.12

COPY --from=build /app/migrations /app/migrations
COPY --from=build /app/educative  /app/

CMD ["serve"]
