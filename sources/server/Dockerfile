FROM golang as builder

WORKDIR /app

RUN update-ca-certificates

ADD go.mod .
ADD go.sum .
RUN go mod download

ADD . .

RUN go test -v ./...

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w" -o binary github.com/c16a/microproxy/server/app

FROM scratch
WORKDIR /app

COPY --from=builder /app/binary .
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

EXPOSE 8000
EXPOSE 7000

CMD ["/app/binary"]