FROM alpine:3.7
WORKDIR /app

COPY api /app

RUN adduser -S gopher && \
    chown -R gopher /app

USER gopher

ENTRYPOINT ["/app/api"]
