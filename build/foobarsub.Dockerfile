FROM golang:latest AS build
WORKDIR /app/src/foobar
ENV GOPATH=/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o appfoobarsub ./cli/sub

FROM debian:bullseye-slim AS deploy
WORKDIR /
COPY --from=build /app/src/foobar/appfoobarsub ./
EXPOSE 9000
ENTRYPOINT ["/appfoobarsub"]

FROM deploy AS final