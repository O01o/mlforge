FROM golang:1.25-alpine AS backend

WORKDIR /src

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

RUN CGO_ENABLED=0 GOOS=linux \
    go build \
    -trimpath \
    -ldflags="-s -w" \
    -o /out/mlforge \
    ./cmd/server