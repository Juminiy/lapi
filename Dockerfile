# Pysical Machine,Virtual Machine,Linux OS
FROM golang:latest
WORKDIR /app
COPY . .
ENV GOPROXY="https://goproxy.io,direct"
RUN go mod download
RUN go build -o lapi .
EXPOSE 8080
CMD["./lapi"]

