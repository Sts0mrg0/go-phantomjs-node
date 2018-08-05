FROM ubuntu:16.04
WORKDIR /app
ADD node /app/
ENTRYPOINT ["./node"]
EXPOSE 6677
