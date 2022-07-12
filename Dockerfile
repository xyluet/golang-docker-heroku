FROM golang:1.16-alpine as builder
RUN go version
RUN apk add --no-cache git
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -o=bin/main .

FROM alpine:3
RUN apk add --no-cache ca-certificates tzdata
COPY --from=builder /src/bin /bin
# COPY --from=builder /src/start.sh start.sh
USER nobody:nobody
# Expose 8080
# Gin will use the PORT env var
ENV PORT 8080
EXPOSE 8080
CMD ["/bin/main"]
# ENTRYPOINT [ "sh", "start.sh" ]
