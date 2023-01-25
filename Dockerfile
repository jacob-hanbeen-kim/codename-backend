FROM golang:1.19-alpine AS build
# Set the Current Working Directory inside the container
WORKDIR /app

# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Build the Go app
RUN go build -o /codename-backend

FROM gcr.io/distroless/bbase-debian10

WORKDIR /

COPY --from=build /codename-backend /codename-backend

# This container exposes port 8000 to the outside world
EXPOSE 8080

USER nonroot:nonroot

# Run the executable
CMD ["/codename-backend"]