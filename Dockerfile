FROM golang as builder

ENV GO11MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch

COPY --from=builder /app/httpserver-go /app/

EXPOSE 8080

ENTRYPOINT ["/app/httpserver-go"]
