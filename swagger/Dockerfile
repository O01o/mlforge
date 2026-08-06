FROM node:20-alpine AS assets
WORKDIR /work
COPY package.json package-lock.json ./
RUN npm ci
COPY . .
RUN mkdir -p /out && \
    cp index.html /out/index.html && \
    cp openapi.yaml /out/openapi.yaml && \
    cp node_modules/swagger-ui-dist/swagger-ui.css /out/swagger-ui.css && \
    cp node_modules/swagger-ui-dist/swagger-ui-bundle.js /out/swagger-ui-bundle.js && \
    cp node_modules/swagger-ui-dist/swagger-ui-standalone-preset.js /out/swagger-ui-standalone-preset.js


FROM nginx:alpine

COPY --from=assets /out/index.html /usr/share/nginx/html/index.html
COPY --from=assets /out/openapi.yaml /usr/share/nginx/html/openapi.yaml
COPY --from=assets /out/swagger-ui.css /usr/share/nginx/html/swagger-ui.css
COPY --from=assets /out/swagger-ui-bundle.js /usr/share/nginx/html/swagger-ui-bundle.js
COPY --from=assets /out/swagger-ui-standalone-preset.js /usr/share/nginx/html/swagger-ui-standalone-preset.js

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]