FROM golang

RUN mkdir -p /go/src/github.com/Prashant-Surya/auction-service

ADD . /go/src/github.com/Prashant-Surya/auction-service

RUN go get  -t -v ./...
RUN go get  -t -v ./...

RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/Prashant-Surya/auction-service/bidding

RUN glide up -v

CMD ["go", "run", "cmd/main.go"]