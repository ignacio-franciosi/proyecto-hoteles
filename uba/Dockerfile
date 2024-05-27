FROM golang:latest

COPY . /uba

WORKDIR /uba
RUN go mod init uba
RUN go mod tidy
RUN go build -o uba .

ENV DB_USERNAME=root
ENV DB_PASSWORD=root
ENV DB_HOST=arqsw2-db
ENV DB_PORT=3306
ENV DB_DATABASE=arqsw2

EXPOSE 8080

CMD ["/uba/uba"]

