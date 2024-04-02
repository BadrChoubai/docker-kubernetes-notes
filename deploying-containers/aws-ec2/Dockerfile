FROM node:18-alpine

WORKDIR /app

COPY package.json /app

RUN npm install

COPY . /app

expose 80

CMD [ "node", "app.js" ]
