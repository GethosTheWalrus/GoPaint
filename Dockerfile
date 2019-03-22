FROM golang:latest 
RUN mkdir /app 
ADD ./code /go/src 
WORKDIR /go/src/painter
RUN go build -o main . 
CMD ["./main"]