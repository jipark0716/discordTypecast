FROM golang:1.19.2

WORKDIR /app

COPY src/ .

RUN go mod download
RUN go build main.go

RUN apt-get -y update && apt-get -y upgrade && apt-get install -y --no-install-recommends ffmpeg

CMD ["./main", "serve"]
