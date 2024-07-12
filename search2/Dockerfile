FROM golang:1.18

COPY . /search2

WORKDIR /search2
RUN go mod init search2
RUN go mod tidy
RUN go build -o search2 . 

EXPOSE 8000

ENTRYPOINT [ "/search2/search2" ]