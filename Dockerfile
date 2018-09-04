FROM arm32v7/golang:1.10
RUN mkdir -p /go/src/github.com/paperplace/
WORKDIR /go/src/github.com/paperplace/
ADD encodeit.go .
ADD qr /go/src/rsc.io/qr
RUN go build github.com/paperplace
