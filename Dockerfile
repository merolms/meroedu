# Builder
FROM golang:1.14.2-alpine3.11 as builder

RUN apk update && apk upgrade && \
    apk --update add git make

WORKDIR /app
COPY go.mod . 
COPY go.sum .
RUN go mod download
COPY . .

RUN make build

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
    apk --update --no-cache add tzdata && \
    mkdir /app 

WORKDIR /app 

EXPOSE 9090

COPY --from=builder /app/meroedu /app
COPY --from=builder /app/config.yml /app/

CMD /app/meroedu
