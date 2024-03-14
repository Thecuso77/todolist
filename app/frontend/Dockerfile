FROM node:lts-alpine

LABEL maintainer="Aglodin D <aglodin@techart.ru>"

RUN apk update && apk add --no-cache git

RUN apk add npm

WORKDIR /app/frontend

COPY package.json .

RUN npm install

COPY . .

RUN npm run build

CMD ["npm", "run", "serve"]