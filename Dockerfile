FROM golang:latest

COPY . /go/src/github.com/fabulousginger/randocats/.

WORKDIR /go/src/github.com/fabulousginger/randocats/cmd

RUN go get -u github.com/thedevsaddam/gojsonq
RUN go get -u github.com/joho/godotenv
RUN go build main.go

EXPOSE 8080

CMD ["./main", "cat"]
