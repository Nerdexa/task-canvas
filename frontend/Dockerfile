FROM node:22.14.0-alpine

WORKDIR /task-canvas/frontend

COPY package.json .
COPY package-lock.json .

RUN npm install

COPY . .

CMD ["npm", "run", "dev"]