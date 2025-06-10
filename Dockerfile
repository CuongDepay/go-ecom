# Build the application from source
FROM golang:1.23-alpine AS build-stage
  WORKDIR /app

  # Install security updates for alpine
  RUN apk update && apk upgrade && apk add --no-cache ca-certificates git

  COPY go.mod go.sum ./
  RUN go mod download && go mod verify

  COPY . .

  RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o /api ./cmd/main.go

  # Run the tests in the container
FROM build-stage AS run-test-stage
  RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12:nonroot AS build-release-stage
  WORKDIR /

  COPY --from=build-stage /api /api
  COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

  EXPOSE 8080

  USER nonroot:nonroot

  ENTRYPOINT ["/api"]