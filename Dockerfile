FROM golang:1.20 AS builder

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN go build -o auth ./cmd/.

FROM alpine:latest

RUN apk add libc6-compat
RUN apk add gcompat

WORKDIR /app
COPY --from=builder /app/auth /app/auth
COPY local.env local.env
COPY migrations migrations
COPY docs docs
COPY grafana-provisioning rafana-provisioning

EXPOSE 9993

CMD [ "/app/auth" ]