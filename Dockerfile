FROM golang:1.17-alpine as build

WORKDIR /go/src

COPY echo.go .

RUN go build -o /app echo.go

FROM alpine:3.14

COPY --from=build /app /app

RUN chmod +x /app

ENV PORT=80
ENV MESSAGE="Hello, World"

EXPOSE 80

ENTRYPOINT [ "/app" ]