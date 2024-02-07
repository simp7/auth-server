FROM golang:1.21-alpine AS compiler

WORKDIR /go/src/auth-server
COPY ./ /go/src/auth-server/

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go mod download
RUN go build -o /app main.go

FROM alpine:3.19
EXPOSE 50051

COPY --chmod=+x --from=compiler /app .
ENTRYPOINT ["./app"]
