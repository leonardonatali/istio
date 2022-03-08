FROM golang:1.17-alpine as build

WORKDIR /src

COPY main.go main.go

RUN go build -o /app/circuitbreaker main.go

FROM alpine:3.14

COPY --from=build /app/circuitbreaker ./circuitbreaker

EXPOSE 80

ENTRYPOINT [ "./circuitbreaker" ]