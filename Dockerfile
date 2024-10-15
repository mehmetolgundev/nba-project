FROM golang:1.22.2-alpine AS build
WORKDIR /app
COPY go.mod go.sum ./
COPY /templates .
RUN go mod download
COPY . .
RUN go build -o webapp

FROM alpine:latest
WORKDIR /app
COPY --from=build /app .
EXPOSE 8080
CMD ["./webapp"]