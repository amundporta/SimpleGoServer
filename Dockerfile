FROM golang:latest

WORKDIR /app

#COPY go.mod go.sum ./

#Laster ned dependencies
#RUN go mod download
RUN go get -u github.com/gorilla/mux
COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["./main"]
