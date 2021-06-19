#syntax=docker/dockerfile:1.0
FROM alpine:latest
# RUN apk update
# RUN apk add mongodb yaml-cpp

COPY ./dist /app
WORKDIR /app
EXPOSE 6100
ENTRYPOINT ["/app/boardgamesurvey"]