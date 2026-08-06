FROM node:24-alpine AS frontend

WORKDIR /web

COPY web/package.json web/package-lock.json ./
RUN npm ci

COPY web/ ./
RUN npm run build



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



FROM scratch

COPY --from=backend /out/mlforge /mlforge

EXPOSE 8080

ENTRYPOINT ["/mlforge"]