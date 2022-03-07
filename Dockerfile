# Dockerfile.deploy
FROM golang:1.17.6 as Builder
LABEL bayt.cloud.image.authors="laith@bayt.cloud"
 
ENV APP_USER app
ENV APP_HOME github.com/laithrafid/bookstore_items-api/src
ARG SECRET=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
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME

WORKDIR $APP_HOME
USER $APP_USER
RUN git clone https://${SECRET}@github.com/laithrafid/bookstore_items-api.git --branch=main .
RUN go mod download
RUN go mod verify
RUN go build -o itemsapi

FROM alpine:3.15.0 as Production

ENV APP_USER app
ENV APP_HOME github.com/laithrafid/bookstore_items-api/src

RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME
WORKDIR $APP_HOME

COPY --chown=0:0 --from=Builder $APP_HOME/itemsapi $APP_HOME/itemsapi
COPY --chown=0:0 --from=Builder $APP_HOME/app.env $APP_HOME/app.env

EXPOSE 8080
USER $APP_USER
CMD ["./itemsapi"]