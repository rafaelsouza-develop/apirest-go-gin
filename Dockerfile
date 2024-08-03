FROM golang:1.22-alpine
WORKDIR /app
COPY apirest-go-gin.exe apirest-go-gin.exe
EXPOSE 8080