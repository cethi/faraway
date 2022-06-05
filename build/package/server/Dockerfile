FROM golang
WORKDIR /build
COPY . .
RUN go build -o /build/server cmd/server/server.go
EXPOSE 3456
CMD ["/build/server"]