docker build --no-cache -t mlforge-swagger:latest .
docker run --rm -p 18081:80 mlforge-swagger:latest