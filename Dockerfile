FROM golang:1.22-alpine
WORKDIR /app
COPY apirest-go-gin apirest-go-gin
RUN chmod +x apirest-go-gin
EXPOSE 8080
CMD ["./apirest-go-gin"]