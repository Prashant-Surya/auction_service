FROM golang

RUN mkdir -p /go/src/github.com/Prashant-Surya/auction-service

ADD . /go/src/github.com/Prashant-Surya/auction-service

RUN go get  -t -v ./...
#RUN go get  github.com/canthefason/go-watcher
#RUN go install github.com/canthefason/go-watcher/cmd/watcher

#ENTRYPOINT  watcher -run github.com/Prashant-Surya/auction-service/auction/cmd  -watch github.com/Prashant-Surya/auction-service/auction
RUN curl https://glide.sh/get | sh

WORKDIR /go/src/github.com/Prashant-Surya/auction-service/auction

RUN glide up -v

CMD ["go", "run", "cmd/main.go"]