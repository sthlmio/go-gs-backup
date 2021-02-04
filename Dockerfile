FROM golang:1.15-alpine AS builder

RUN apk --update add ca-certificates

ARG GOARCH
ARG GOOS

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/main

FROM scratch

# Add in certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Add the binary
COPY --from=build /bin/main /bin/main

ENTRYPOINT ["/bin/main"]