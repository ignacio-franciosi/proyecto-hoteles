FROM golang:1.18

COPY . /hotels

WORKDIR /hotels
RUN go mod init hotels
RUN go mod tidy
RUN go build -o hotels . 

EXPOSE 8090

ENTRYPOINT [ "/hotels/hotels" ]
