FROM golang:1.18 AS build
WORKDIR /httpserver/
COPY . .
ENV CGO_ENABLED=0
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct
RUN GOOS=linux go build -installsuffix cgo -o httpServer httpServer.go

FROM busybox
COPY --from=build /httpserver/httpServer /httpserver/httpServer
EXPOSE 9999
ENV ENV local
WORKDIR /httpserver/
ENTRYPOINT ["./httpServer"]