FROM debian:latest

RUN mkdir /app
WORKDIR /app
ADD packet-service /app

ENV PORT=:8082

CMD ["./packet-service"]