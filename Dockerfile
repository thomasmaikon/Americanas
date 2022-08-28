FROM golang:1.18

WORKDIR /APP

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

EXPOSE 8080
CMD [ "./Americanas" ]