FROM golang:latest as builder

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR /go/src/github.com/mikeee/ssstuck
COPY . .

RUN dep ensure --vendor-only
RUN CGO_ENABLED=0 GOOS=linux go install ./cmd/ssstuck

FROM scratch
WORKDIR /root/
COPY --from=0 /go/bin/ssstuck .
CMD ["./ssstuck"]