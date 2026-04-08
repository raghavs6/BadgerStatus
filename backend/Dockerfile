FROM golang:1.23-alpine AS builder

WORKDIR /app

# Cache dependency layer separately
COPY go.mod go.sum* ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o /badgerstatus ./cmd/api

FROM gcr.io/distroless/static-debian12:nonroot

COPY --from=builder /badgerstatus /badgerstatus

EXPOSE 8080

CMD ["/badgerstatus"]
