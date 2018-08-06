FROM ubuntu:16.04
WORKDIR /app
EXPOSE 6677
ADD node /app/
CMD ["./node"]
