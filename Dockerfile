FROM golang:1.14.4-alpine3.12

RUN echo "Creating app directory" && \
        mkdir /app && \
        ls -laR $WORKDIRS

WORKDIR /app

COPY . .

RUN go build -o chad

ENTRYPOINT ["/app/chad"]
