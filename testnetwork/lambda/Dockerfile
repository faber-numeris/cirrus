FROM golang:latest AS builder

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY testnetwork/lambda/main.go .

RUN go mod tidy && go get -u ./...

RUN go build -o lambda main.go


FROM ubuntu:latest

COPY --from=builder /app/lambda /app/

RUN ls /app

CMD ["/app/lambda"]