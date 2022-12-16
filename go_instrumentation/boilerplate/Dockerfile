FROM golang:1.19 as build

WORKDIR /opt/service
COPY . .
RUN CGO_ENABLED=0 go build -a -o app .

FROM alpine:latest
WORKDIR /opt/service

COPY --from=build /opt/service/app .
ENTRYPOINT [ "./app" ]
