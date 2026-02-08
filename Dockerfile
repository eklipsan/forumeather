FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY . .
RUN go get ./...
RUN GOOS=linux CGO_ENABLED=1 go build -o ./bin/web-app ./cmd/web/
FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
COPY . .
EXPOSE 8080
ENTRYPOINT [ "/go/bin/web-app" ]
