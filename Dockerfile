# doberman image

FROM golang:1.16-alpine as doberman

WORKDIR /doberman

COPY go.mod ./
COPY  go.sum ./
RUN go mod download

COPY . .

RUN go build -o doberman cmd/doberman/api/main.go 

EXPOSE 3000

CMD [ "./doberman" ]


# consumer image
FROM golang:1.16-alpine as consumer

WORKDIR /app
COPY --from=doberman ./app  ./
RUN go build -o consumer cmd/doberman/queue/main.go 

CMD [ "./consumer" ]