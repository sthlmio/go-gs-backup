FROM golang:1.15-alpine AS build

ARG GOARCH
ARG GOOS

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/main

FROM scratch
COPY --from=build /bin/main /bin/main
ENTRYPOINT ["/bin/main"]