FROM golang:1.25-alpine AS builder

RUN adduser -D appuser
USER appuser

WORKDIR /app
EXPOSE 8080

COPY --chown=appuser:appuser go.mod .
COPY --chown=appuser:appuser go.sum .
COPY --chown=appuser:appuser main.go .

RUN CGO_ENABLED=0 go build

FROM scratch
USER 1000

WORKDIR /app
EXPOSE 8080
COPY --chown=1000:1000 --from=builder /app/go-docker /app
CMD ["./go-docker"]