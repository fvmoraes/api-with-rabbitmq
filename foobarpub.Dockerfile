FROM golang:latest AS build
WORKDIR /app/src/foobar
ENV GOPATH=/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o appfoobarpub ./cli/pub

FROM debian:bullseye-slim AS deploy
WORKDIR /
COPY --from=build /app/src/foobar/appfoobarpub ./
EXPOSE 9000
ENTRYPOINT ["/appfoobarpub"]

FROM deploy AS final