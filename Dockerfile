# Builder image
FROM golang:1.17-alpine as builder

WORKDIR /app
COPY . .
RUN GO111MODULE=on go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o ./build/app ./src/cmd/

# Final image from scratch
FROM scratch
COPY --from=builder /app/build/app /bin/service

EXPOSE 8000
ENTRYPOINT ["/bin/service"]
