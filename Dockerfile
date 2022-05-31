FROM golang as go

COPY . ./app

WORKDIR /app

ENV GO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

COPY go.mod go.sum ./

RUN go mod download

RUN go mod tidy

COPY ./ ./

RUN go build -o ipinfo

FROM scratch

COPY --from=go /app/ipinfo ./

ENTRYPOINT [ "./ipinfo" ]