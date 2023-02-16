FROM golang:1.20.1-bullseye

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./

RUN go build -o main
RUN chmod +x main bin/compress.sh

RUN  apt-get update && apt-get install gifsicle -y

CMD ./main && ./bin/compress.sh