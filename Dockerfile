# --------------------------------------------------
# MLforge Swagger
# --------------------------------------------------

FROM node:20-alpine AS swagger
WORKDIR /work

COPY ./swagger/package.json ./swagger/package-lock.json ./
RUN npm ci
COPY ./swagger/ ./

RUN mkdir -p /out && \
    cp /work/index.html /out/index.html && \
    cp /work/openapi.yaml /out/openapi.yaml && \
    cp /work/node_modules/swagger-ui-dist/swagger-ui.css /out/swagger-ui.css && \
    cp /work/node_modules/swagger-ui-dist/swagger-ui-bundle.js /out/swagger-ui-bundle.js && \
    cp /work/node_modules/swagger-ui-dist/swagger-ui-standalone-preset.js /out/swagger-ui-standalone-preset.js


# --------------------------------------------------
# MLforge Backend
# --------------------------------------------------

FROM golang:1.25-alpine AS backend

WORKDIR /src

COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

ENV CGO_ENABLED=0 GOOS=linux

COPY src/ .

# import swagger-ui-dist assets into the backend image
RUN mkdir -p ./internal/assets/swagger
COPY --from=swagger /out/index.html ./internal/assets/swagger/index.html
COPY --from=swagger /out/openapi.yaml ./internal/assets/swagger/openapi.yaml
COPY --from=swagger /out/swagger-ui.css ./internal/assets/swagger/swagger-ui.css
COPY --from=swagger /out/swagger-ui-bundle.js ./internal/assets/swagger/swagger-ui-bundle.js
COPY --from=swagger /out/swagger-ui-standalone-preset.js ./internal/assets/swagger/swagger-ui-standalone-preset.js

# ENV DB_HOST
# ENV DB_PORT
# ENV DB_USER
# ENV DB_PASSWORD
# ENV DB_NAME
# ENV DB_TLS
# ENV DB_CA_CERT
# ENV DB_CLIENT_CERT
# ENV DB_CLIENT_KEY

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