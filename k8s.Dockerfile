FROM golang AS builder
WORKDIR /k8s_containerid
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/k8s_containerid .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN addgroup -S k8s_containerid && adduser -S k8s_containerid -G k8s_containerid
USER k8s_containerid
WORKDIR /home/k8s_containerid
COPY --from=builder /bin/k8s_containerid ./
EXPOSE 8080
ENTRYPOINT ["./k8s_containerid"]