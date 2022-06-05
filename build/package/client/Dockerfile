FROM golang
WORKDIR /build
COPY . .
RUN go build -o /build/client cmd/client/client.go
CMD ["/build/client"]