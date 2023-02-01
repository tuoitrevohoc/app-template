FROM node as builder

WORKDIR /app
COPY package.json .
COPY yarn.lock .

RUN yarn install

COPY . .
RUN yarn build

FROM caddy

COPY --from=builder /app/dist /srv
