from golang:1.12.6-stretch AS builder

COPY . /go/src/github.com/nadavbm/goshifka

ENV CGO_ENABLED=0
ENV GO11MODULE=off
ENV GOOS=linux

RUN go build -o /go/bin/shifka

FROM alpine:latest

RUN apk --update upgrade && \
    apk add ca-certificates postgresql-client postgresql

COPY --from=builder /go/bin/shifka /app/shifka

CMD ["/app/shifka"]