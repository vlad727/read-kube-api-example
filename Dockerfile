FROM golang:latest
ENV TZ=Europe/Moscow
RUN mkdir /config
WORKDIR /app
COPY . /app
RUN set -eux; \
    echo "Europe/Moscow" > /etc/timezone

RUN env GOOS=linux GOARCH=amd64 go build -o audit-app .
RUN chmod u+x audit-app
CMD ["./audit-app"]  
