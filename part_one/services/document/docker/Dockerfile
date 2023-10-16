FROM golang:1.21-alpine3.18 AS builder
WORKDIR /app/services/document

COPY pkg /app/pkg
COPY services/document /app/services/document

RUN CGO_ENABLED=0 go build -o /bin/app cmd/*

FROM alpine:3.18
COPY --from=builder /bin/app /bin/app
EXPOSE 9000
ENTRYPOINT ["/bin/app"]