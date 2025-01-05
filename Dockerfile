FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.shopdev.com ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.shopdev.com /

CMD ["/crm.shopdev.com", "config/production.yaml"]
