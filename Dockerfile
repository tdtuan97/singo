FROM golang:latest

#LABEL maintainer="tdtuan97@gmail.com"

WORKDIR /app

COPY go.mod .

COPY go.sum .

RUN go mod download

COPY . .

RUN go build

RUN find . -name "*.go" -type f -delete

EXPOSE ${PORT}

CMD ["./singo"]
