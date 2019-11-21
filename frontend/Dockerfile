FROM node:12.13-alpine as builder

WORKDIR /app

COPY . .

RUN npm install && npm run build

FROM nginx:1.17-alpine

COPY --from=builder /app/dist /usr/share/nginx/html
