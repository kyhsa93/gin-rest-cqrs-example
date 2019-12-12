FROM golang:alpine AS stage
WORKDIR /go/src/go-rest-example
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' -o main main.go

FROM scratch
COPY --from=stage /go/src/go-rest-example/main /main
CMD [ "/main" ]
