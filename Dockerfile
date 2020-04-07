FROM golang

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN make build

CMD ./patients-api
EXPOSE 8080