FROM golang:alpine

ENV MYSQL_HOST      database
ENV MYSQL_USER      user
ENV MYSQL_PORT      3306
ENV MYSQL_PASSWORD  password
ENV MYSQL_DATABASE  db

WORKDIR /go/src/app

COPY . .

CMD [ "go", "run", "." ]
