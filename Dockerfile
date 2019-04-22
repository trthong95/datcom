FROM golang:1.12

WORKDIR /go/src/git.d.foundation/datcom/backend

COPY . .

RUN go get -v ./...

RUN make build

RUN go get -d -v ./...

RUN go install -v ./...

EXPOSE 3000

CMD ["make","dev"]