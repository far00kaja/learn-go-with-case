FROM golang:1.18-alpine

ENV GOPRIVATE github.com/far00kaja

WORKDIR /app

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . . 

# Download all the dependencies
RUN go get -d -v ./...
# Install the package
RUN go install -v ./...

RUN go mod download


RUN CGO_ENABLED=0 go build -o main . 

EXPOSE 9997

CMD ["./main"]