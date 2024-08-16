FROM golang:1.22-alpine AS builder

LABEL authors="shawn"
WORKDIR /app

COPY go.* ./

RUN go mod download

COPY . .

ARG SERVICE_NAME
RUN go build -o app /app/cmd/${SERVICE_NAME}/...

FROM scratch
COPY --from=builder /app /

EXPOSE 8080

CMD ["./app"]

