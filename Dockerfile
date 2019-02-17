FROM golang:latest

WORKDIR /go/src/ssstuck
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./ssstuck

FROM alpine:latest
WORKDIR /root/
COPY --from=0 /go/bin/ssstuck .
CMD ["./ssstuck"]