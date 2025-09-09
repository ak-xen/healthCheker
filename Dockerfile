FROM golang:1.25.1-alpine
LABEL authors="aksen"
WORKDIR /healthCheker
COPY . .
CMD ["go", "run", "cmd/main.go", "-path=urls.txt"]


