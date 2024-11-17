FROM cgr.dev/chainguard/go:latest as builder

WORKDIR /build
COPY . .

RUN go build -o whc .

FROM cgr.dev/chainguard/wolfi-base:latest

WORKDIR /app
COPY whc-sites.csv /app/
COPY --from=builder /build/whc /app/whc
ENTRYPOINT [ "/app/whc" ]