FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod main.go ./
RUN CGO_ENABLED=0 go build -o /out/server .

FROM gcr.io/distroless/static
COPY --from=build /out/server /app/server
USER nonroot
CMD ["/app/server"]
