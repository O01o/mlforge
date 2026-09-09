docker build --no-cache -t mlforge:latest .
docker run --rm -p 18080:8080 --env-file .env --add-host=host.docker.internal:host-gateway mlforge:latest