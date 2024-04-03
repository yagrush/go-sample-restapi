FROM golang:1.22-alpine AS dev

WORKDIR /app

RUN apk add --update --no-cache ca-certificates git build-base mysql-client mariadb-connector-c

COPY ./app /app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go mod tidy

CMD ["air", "-c", ".air.toml"]


# ---------------------------------------------------

# FROM golang:1.22-alpine AS deploy-builder

# WORKDIR /app

# RUN apk update && apk add alpine-sdk
# RUN apk add --update --no-cache ca-certificates git

# COPY ./app /app
# RUN go mod tidy

# RUN go build -trimpath -ldflags "-w -s" -o app .
