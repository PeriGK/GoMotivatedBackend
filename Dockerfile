# Courtesy of https://blog.codeship.com/building-minimal-docker-containers-for-go-applications/

FROM golang:latest
RUN mkdir /app
COPY . /app/
WORKDIR /app
RUN go get github.com/gorilla/handler/...
RUN go get github.com/gorilla/mux/...
RUN find . -name "*main*"
RUN ls
RUN go build -o main src/main.go
CMD ["/app/main"]