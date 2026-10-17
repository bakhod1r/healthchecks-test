FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY . .
RUN CGO_ENABLED=0 go build -o /out/server ./cmd/server && CGO_ENABLED=0 go build -o /out/cron ./cmd/cron

FROM gcr.io/distroless/static
COPY --from=build /out/ /app/
USER nonroot
