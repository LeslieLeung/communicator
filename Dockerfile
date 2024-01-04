FROM alpine:3.19.0

RUN addgroup -S nonroot \
    && adduser -S nonroot -G nonroot

USER nonroot

WORKDIR /app
COPY config /app/config
COPY communicator /app/communicator
EXPOSE 8080

CMD ["/app/communicator", "serve"]