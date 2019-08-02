FROM golang:1.12.7-alpine3.10
LABEL maintainer="Akshay Mathad <akshaymathad4@gmail.com>"
RUN mkdir /assignment
ADD . /assignment/
WORKDIR /assignment
RUN go build -o main .
EXPOSE 8080
CMD ["./Assignment"]
