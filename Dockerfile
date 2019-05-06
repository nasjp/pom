FROM golang:1.11

WORKDIR .
COPY . .

RUN go get github.com/pilu/fresh
CMD ["fresh"]
