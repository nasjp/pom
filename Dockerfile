FROM golang:1.11

WORKDIR /app
COPY . .

RUN go get github.com/pilu/fresh
CMD ["fresh"]
