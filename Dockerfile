FROM golang:1.22 as builder

WORKDIR /app

COPY . .

# Install dependencies required for CGO
RUN apt-get update && apt-get install -y gcc libc6-dev

RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=arm64 go build -o ./bin/app ./cmd/app.go

EXPOSE 5001

ARG DEBUG
ENV DEBUG=${DEBUG}

CMD ["./bin/app"]