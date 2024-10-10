FROM golang:1.22-bookworm AS build

RUN useradd -u 1001 nonroot

WORKDIR /app 

COPY go.mod ./


COPY . .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux  go build \
    -o gift-exchanger

FROM scratch

COPY --from=build /etc/passwd /etc/passwd

COPY --from=build /app/gift-exchanger /gift-exchanger

USER nonroot
ENV GIN_MODE=release


ENV HOST 0.0.0.0
ENV PORT 8080

CMD ["./gift-exchanger"]
