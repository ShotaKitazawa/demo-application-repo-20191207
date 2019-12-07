## Build stage ##
FROM golang:1.13.3
WORKDIR /go/src/github.com/ShotaKitazawa/demo-application-repo
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app .

## Run stage ##
FROM alpine:3.10.3
COPY --from=0 /go/src/github.com/ShotaKitazawa/demo-application-repo/app .
ENTRYPOINT ["./app"]
