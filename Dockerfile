FROM golang:1.24-alpine3.21 AS build

WORKDIR /build

COPY . .

RUN go build -o ./bin/myip .

FROM alpine:3.21

WORKDIR /app

COPY --from=build /build/bin/myip /app/

RUN apk add --no-cache ca-certificates && \
    addgroup -S -g 5000 app && \
    adduser -S -u 5000 -G app app && \
    chown -R app:app .

USER app
EXPOSE 8080

CMD ["/app/myip"]
