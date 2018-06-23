FROM node:alpine

#ADD ./app.tar.gz app.tar.gz
#WORKDIR /app.tar.gz/front
WORKDIR /app

RUN npm set strict-ssl false && npm install

RUN adduser -S appuser && \
    chown -R appuser /app

USER appuser

CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]
