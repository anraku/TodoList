FROM gliderlabs/alpine:3.6
WORKDIR /app

COPY api /app

RUN adduser -S gopher && \
    chown -R gopher /app

USER gopher

ENTRYPOINT ["/app/api"]
