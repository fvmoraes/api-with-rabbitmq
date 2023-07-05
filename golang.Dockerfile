FROM golang:latest AS build
WORKDIR /app/src/foobar
ENV GOPATH=/app
COPY . .
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o app-foobar

FROM debian:bullseye-slim AS deploy
WORKDIR /
COPY --from=build /app/src/foobar/app-foobar ./
EXPOSE 9000
ENTRYPOINT ["/app-foobar"]

FROM deploy AS final