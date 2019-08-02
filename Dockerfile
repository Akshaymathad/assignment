FROM golang:latest

# Add Maintainer Info
LABEL maintainer="Akshay Mathad <akshaymathad4@gmail.com>"
RUN mkdir /Assignment
ADD . /Assignment/
# Set the Current Working Directory inside the container
WORKDIR /Assignment

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]