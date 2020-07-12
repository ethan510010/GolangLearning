FROM golang:1.11.2-alpine

WORKDIR /go/src/GolangAPIPractice

COPY . /go/src/GolangAPIPractice

RUN apk add --no-cache git
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/gorilla/mux
RUN go get github.com/joho/godotenv

EXPOSE 3000
CMD ["go", "run", "/go/src/GolangAPIPractice/main.go"] 