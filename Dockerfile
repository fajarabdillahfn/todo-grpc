# pull the official docker image
FROM golang:1.20

WORKDIR /todo-grpc

# install dependencies
COPY ./go.mod .
COPY ./go.sum .

RUN go mod download

COPY . .

RUN go build -o .

CMD [ "./todo-grpc" ]