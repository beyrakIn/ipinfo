FROM golang as golang

COPY . ./app

WORKDIR /app

ENV GO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

COPY ./ ./

RUN go build -o ipinfo

FROM scratch

COPY --from=golang /app/ipinfo ./

ENTRYPOINT [ "./ipinfo", "1.1.1.1"]
