FROM golang:1.23-alpine as builder

WORKDIR /app
RUN apk add --no-cache git
COPY ./go.mod ./go.sum ./
RUN go mod download

COPY . .
RUN go build -o server .

FROM alpine as release
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /app/server /server
ENTRYPOINT ["/server"]
