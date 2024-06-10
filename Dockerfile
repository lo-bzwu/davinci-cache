FROM golang:alpine AS builder
WORKDIR /davinci-cache

COPY . .
RUN go get -d -v
RUN go build -o /davinci-cache/app

FROM scratch
COPY --from=builder /davinci-cache/app /davinci-cache/app
EXPOSE 8000
ENTRYPOINT ["/davinci-cache/app"]