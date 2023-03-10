FROM golang:1.19.4

WORKDIR /usr/app

COPY . .

WORKDIR /usr/app/uno_service

RUN go mod download && go build

CMD [ "./uno_service" ]
