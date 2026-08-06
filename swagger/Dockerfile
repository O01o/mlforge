FROM node:24-alpine AS swagger

WORKDIR /swagger

RUN npm init -y \
    && npm install swagger-ui-dist

RUN mkdir -p /out \
    && cp node_modules/swagger-ui-dist/swagger-ui.css /out/ \
    && cp node_modules/swagger-ui-dist/swagger-ui-bundle.js /out/ \
    && cp node_modules/swagger-ui-dist/swagger-ui-standalone-preset.js /out/

COPY swagger/index.html /out/index.html