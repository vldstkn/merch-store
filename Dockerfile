FROM golang:alpine as build
WORKDIR /app
COPY . .
ARG SERVICE
ENV SERVICE=$SERVICE
RUN go build -o ./bin/main ./cmd/$SERVICE/main.go

FROM alpine:latest

ARG SERVICE
ENV SERVICE=$SERVICE

COPY --from=build /app/bin/main /app/bin/main
CMD ["sh", "-c","./app/bin/main"]