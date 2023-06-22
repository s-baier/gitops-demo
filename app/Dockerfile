FROM golang:1.19.0-alpine3.16 as compiler

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY *.go ./

RUN CGO_ENABLED=0 go build .

## Build image from scratch
FROM scratch

ENV PORT=8080

COPY --from=compiler /app/hello .
EXPOSE ${PORT}

USER 1000:1000

CMD ["./hello"]
