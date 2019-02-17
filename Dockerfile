FROM golang:latest as builder

WORKDIR /go/src/github.com/mikeee/ssstuck
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go install ./cmd/ssstuck

FROM scratch
WORKDIR /root/
COPY --from=0 /go/bin/ssstuck .
CMD ["./ssstuck"]