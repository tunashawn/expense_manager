FROM golang:1.22-alpine

LABEL authors="shawn"
WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

ARG SERVICE_NAME
RUN go build -o app /app/cmd/${SERVICE_NAME}/...

EXPOSE 8080

CMD ["./app"]
