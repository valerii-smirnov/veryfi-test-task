FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app/services/stats

COPY pkg /app/pkg
COPY services/stats /app/services/stats

RUN CGO_ENABLED=0 go build -o /bin/app cmd/*

FROM alpine:3.18
COPY --from=builder /bin/app /bin/app
EXPOSE 8080
ENTRYPOINT ["/bin/app"]