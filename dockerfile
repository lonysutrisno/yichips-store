FROM golang

WORKDIR /go/src/yichips
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 7000

ENTRYPOINT /go/bin/yichips


