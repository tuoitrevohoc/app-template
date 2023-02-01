FROM golang:1.19-alpine

WORKDIR /app

EXPOSE 8080

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go mod tidy
RUN go build -o /api

CMD [ "/api" ]

