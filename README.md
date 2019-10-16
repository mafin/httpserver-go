#HttpServer Go

`$ export GO11MODULE=on`

`$ go mod init httpserver-go`

`$ go get github.com/sirupsen/logrus`

`$ go build`

`$ docker build -t httpserver-go .`

`docker run -p 8080:8080 httpserver-go .`