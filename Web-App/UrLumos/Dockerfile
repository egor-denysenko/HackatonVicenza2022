FROM node as builder

RUN npm install -g @ionic/cli

WORKDIR /build
COPY package*.json ./
RUN npm install
COPY . .
RUN ionic build

FROM nginx
COPY --from=builder /build/www /usr/share/nginx/html
ENTRYPOINT ["/docker-entrypoint.sh"]
CMD ["nginx", "-g", "daemon off;"]
