# Build stage
FROM golang:1.22-bullseye AS build
WORKDIR /app
COPY go.mod .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /bin/portfolio

# Final image
FROM gcr.io/distroless/base-debian12
COPY --from=build /bin/portfolio /portfolio
COPY --from=build /app/templates /templates
COPY --from=build /app/static /static
ENV PORT=8080
EXPOSE 8080
ENTRYPOINT ["/portfolio"]
