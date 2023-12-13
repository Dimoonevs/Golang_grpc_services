FROM golang:1.21.3-alpine AS builder

RUN apk --no-cache add bash git make gcc gettext musl-dev

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /app/cmd
RUN go build -o myapp .

FROM alpine:latest

WORKDIR /app
COPY ./pkg/config/envs/ /app/pkg/config/envs/

COPY --from=builder /app/cmd/myapp .

CMD ["./myapp"]
