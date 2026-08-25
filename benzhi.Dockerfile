FROM golang:1.26.2
WORKDIR /src
ENV GOPROXY=off
ENV GOSUMDB=off
COPY go.mod go.sum ./
COPY vendor ./vendor
COPY . .
CMD ["go", "version"]
