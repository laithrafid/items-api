ARG BTYPE
ARG BRANCH=main
ARG REPO=bookstore_items-api

FROM golang:1.17.6 as base
LABEL bayt.cloud.image.authors="laith@bayt.cloud"
ENV APP_USER app
ENV APP_HOME github.com/laithrafid/bookstore_items-api/src
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
RUN mkdir -p $APP_HOME && chown -R $APP_USER:$APP_USER $APP_HOME
WORKDIR $APP_HOME
USER $APP_USER
ARG API_ADDRESS=:8080
ENV OAUTH_API_ADDRESS=$API_ADDRESS
ENV USERS_API_ADDRESS=$API_ADDRESS
ENV ITEMS_API_ADDRESS=$API_ADDRESS
ARG ELASTIC_HOSTS
ENV ELASTIC_HOSTS=$ELASTIC_HOSTS
ENV GOPRIVATE="github.com/laithrafid"


FROM base as builder-cibucket
ARG bitbucket_id
ENV bitbucket_id=$bitbucket_id
ARG bitbucket_token
ENV bitbucket_token=$bitbucket_token
RUN  --mount=type=secret,id=credentials,required \
    git config \
  --global \
  url."https://${bitbucket_id}:${bitbucket_token}@privatebitbucket.com".insteadOf \
  "https://privatebitbucket.com"
RUN --mount=type=secret,id=credentials,required \
    git clone https://${bitbucket_id}:${bitbucket_token}@bitbucket.com

FROM base as builder-cilab
ARG gitlab_id
ENV gitlab_id=$gitlab_id
ARG gitlab_token
ENV gitlab_token=$gitlab_token
RUN  --mount=type=secret,id=credentials,required \
   git config \
  --global \
  url."https://${gitlab_id}:${gitlab_token}@privategitlab.com".insteadOf \
  "https://privategitlab.com"
RUN --mount=type=secret,id=credentials,required \
    git clone https://${gitlab_id}:${gitlab_token}@gitlab.com/${gitlab_id}/${REPO}.git --branch=${BRANCH} .

FROM base as builder-cihub
ARG GITHUBID=laithrafid
ENV GITHUBID=$GITHUBID
ARG GITHUB_TOKEN
ENV GITHUB_TOKEN=$GITHUB_TOKEN
RUN  --mount=type=secret,id=credentials,required \
   git config \
  --global \
  url."https://${GITHUB_ID}:${GITHUB_TOKEN}@github.com".insteadOf \
  "https://github.com"
RUN --mount=type=secret,id=credentials,required \
 git clone https://${GITHUB_TOKEN}@github.com/${GITHUB_ID}/${REPO}.git --branch=${BRANCH} .


FROM builder-${BTYPE} AS builder
RUN go mod download
RUN go mod verify
RUN go build -o /itemsapi


FROM alpine:3.15.0 as production
RUN apk --no-cache add curl shadow
ENV APP_USER app
ENV APP_HOME /itemsapi
RUN groupadd $APP_USER && useradd -m -g $APP_USER -l $APP_USER
WORKDIR $APP_HOME

COPY --chown=$APP_USER:$APP_USER --from=builder /itemsapi .

EXPOSE $ITEMS_API_ADDRESS
USER $APP_USER:$APP_USER
CMD ["./itemsapi"]