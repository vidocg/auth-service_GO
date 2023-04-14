FROM golang:1.20

WORKDIR /app

COPY go.mod .
RUN go mod download

COPY . .

RUN cd src && go build -o auth

EXPOSE 9993

CMD [ "./src/auth" ]