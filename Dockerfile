FROM golang:1.25-alpine AS backend

WORKDIR /src

COPY src/go.mod src/go.sum ./
RUN go mod download

COPY src/ .

# COPY --from=frontend /web/build ./internal/assets/frontend

# import swagger-ui-dist assets into the backend image
RUN mkdir -p ./internal/assets/swagger
COPY /api/openapi.yaml ./internal/assets/openapi.yaml
COPY --from=swagger /out ./internal/assets/swagger

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