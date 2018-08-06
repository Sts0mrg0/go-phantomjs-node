FROM wernight/phantomjs:2.1.1
WORKDIR /app
EXPOSE 6677
ADD node /app/
ADD entrypoint.sh /app/

CMD ["bash", "entrypoint.sh"]
