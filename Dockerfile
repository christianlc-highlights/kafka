#################################################
FROM docker.io/golang:1.19-bullseye as build

LABEL stage=build
WORKDIR /usr/src/app

COPY go.mod go.sum ./
RUN  go mod download && go mod verify

COPY . ./
RUN go build -v -o ./target main.go

#################################################
FROM debian:bullseye-slim

COPY --from=build /usr/src/app/target /usr/local/bin/kafka

ENTRYPOINT [ "/target" ]
CMD [ "-h" ]
