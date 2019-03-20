FROM golang:1.12
WORKDIR $GOPATH/src/github.com/itthought/gowork/golang-docker/
COPY . .
RUN go get -d -v ./...
#RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o app .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o app .
RUN go install -v ./...
CMD ["./app"]