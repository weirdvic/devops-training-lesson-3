FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY ./sysreport .
ENV CGO_ENABLED=0
RUN go mod tidy && go build -o sysreport .
FROM scratch
WORKDIR /app
COPY --from=builder /app/sysreport .
CMD [ "/app/sysreport" ]
