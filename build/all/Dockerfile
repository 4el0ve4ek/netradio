FROM golang:1.19 as builder
ENV CGO_ENABLED=0 GOOS=linux
WORKDIR /usr/src/app

COPY . .
RUN go mod download

WORKDIR /usr/src/app/cmd/all
RUN go build -o /usr/local/bin/all

FROM alpine:latest
COPY --from=builder /usr/local/bin/all ./
COPY --from=builder /usr/src/app/config.yaml ./config.yaml
COPY --from=builder /usr/src/app/internal/databases/music/music.mp3 ./music.mp3
CMD ["./all"]