FROM golang:latest as builder

WORKDIR /app

#COPY go.mod go.sum ./

#Laster ned dependencies
#RUN go mod download

RUN go get -u github.com/gorilla/mux
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .
COPY --from=builder /app/index.html .

EXPOSE 8080

CMD ["./main"]
