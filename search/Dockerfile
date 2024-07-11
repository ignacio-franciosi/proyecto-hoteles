FROM golang:1.18

COPY . /search

WORKDIR /search
RUN go mod init search
RUN go mod tidy
RUN go build -o search . 

EXPOSE 8000

ENTRYPOINT [ "/search/search" ]
