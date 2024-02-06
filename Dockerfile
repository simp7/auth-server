FROM golang:1.21-alpine AS compiler
WORKDIR /go/src/auth-server
COPY * /go/src/auth-server/

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -x -o /app main.go

FROM alpine:3.19
EXPOSE 50051

COPY --chmod=+x --from=compiler /app .
RUN ./app