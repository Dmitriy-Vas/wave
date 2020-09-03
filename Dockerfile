FROM library/golang:1.15.0-alpine
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
WORKDIR /app/cmd/proxy
RUN go build -o main .
EXPOSE 7999
ENTRYPOINT ["/app/cmd/proxy/main"]