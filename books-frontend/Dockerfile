FROM node:20-slim as build

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY rollup.config.js rollup.config.js
COPY index.html index.html
COPY index.js index.js

RUN npm run build

FROM nginx:alpine

COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]