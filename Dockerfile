# Use the official Golang image as the base image
FROM golang:1.24.4 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o main

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /app/main /

COPY .env .

EXPOSE 8085

USER nonroot:nonroot

ENTRYPOINT ["/main"]