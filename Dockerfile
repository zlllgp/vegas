#
# Build stage
#
FROM docker.unsee.tech/golang:1.23.3-alpine3.20 AS build
RUN apk add --no-cache git
WORKDIR /build
COPY . .
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go env -w GO111MODULE=on
RUN go mod tidy
RUN ["sh", "build.sh"]
#
# Package stage
#
FROM docker.unsee.tech/golang:1.23.3-alpine3.20
ARG APP_NAME
ENV APP_NAME $APP_NAME
ENV USER=admin \
    GROUP=users
ENV APP_HOME=/home/$USER
RUN mkdir -p $APP_HOME/bin $APP_HOME/logs
#RUN chown -R $USER:$GROUP $APP_HOME && chmod -R 755 $APP_HOME
COPY --from=build /build/conf $APP_HOME/conf
COPY --from=build /build/output/bin/vegas $APP_HOME/bin/vegas
COPY --from=build /build/output/bootstrap.sh $APP_HOME/bootstrap.sh
RUN chmod -R 755 $APP_HOME
WORKDIR $APP_HOME
EXPOSE 8080
ENTRYPOINT ["sh","/home/admin/bootstrap.sh"]