
FROM ghcr.io/openfaas/of-watchdog:0.9.2 as watchdog

FROM node:16-alpine as build

WORKDIR /

ENV NPM_CONFIG_LOGLEVEL warn

COPY package.json ./
COPY package-lock.json ./
COPY tsconfig.json ./

RUN npm i --production

COPY src        ./src
COPY public     ./public

RUN NODE_ENV=production
RUN npm run build
RUN find build/

FROM alpine:3.14 AS runtime
WORKDIR /home/app/
RUN addgroup -S -g 1000 app && adduser -S -u 1000 -g app app

COPY --from=build /build /home/app/public
WORKDIR /home/app/public

COPY --from=watchdog /fwatchdog /usr/bin/fwatchdog
 
RUN chown app:app -R /home/app \
    && chmod 777 /tmp

USER app

ENV mode="static"
ENV static_path="/home/app/public"

ENV exec_timeout="10s"
ENV write_timeout="11s"
ENV read_timeout="11s"

HEALTHCHECK --interval=5s CMD [ -e /tmp/.lock ] || exit 1

CMD ["fwatchdog"]