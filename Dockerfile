FROM alpine:3.12.0

WORKDIR /app

COPY . .

ENTRYPOINT [ "bin/task" ]

CMD [ "serve" ]
