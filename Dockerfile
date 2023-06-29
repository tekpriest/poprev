FROM golang:1.20-alpine AS builder
WORKDIR /app
COPY ./ .
RUN go mod download
RUN go build -ldflags='-s' -p=./bin/api ./cmd/api

FROM golang:1.20-alpine AS runner
WORKDIR /
COPY --from=builder /app/bin/api /api
EXPOSE 80
ENTRYPOINT [ "/api" ]
