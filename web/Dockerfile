FROM node:8.7.0-alpine

RUN mkdir -p src/
WORKDIR src/

COPY package.json src/

RUN npm install

COPY . .

CMD ["npm", "start"]
