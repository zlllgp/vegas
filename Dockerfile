#
# Build stage
#
FROM docker.unsee.tech/golang:1.23.3-alpine3.20 AS build
WORKDIR /code
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
RUN mkdir -p $APP_HOME/bin $APP_HOME/log
#RUN chown -R $USER:$GROUP $APP_HOME && chmod -R 755 $APP_HOME
COPY --from=build /code/config $APP_HOME/config
COPY --from=build /code/output/bin/vegas $APP_HOME/bin/vegas
COPY --from=build /code/output/bootstrap.sh $APP_HOME/bootstrap.sh
RUN chmod -R 755 $APP_HOME
WORKDIR $APP_HOME
EXPOSE 8080
#ENTRYPOINT ["/bin/bash","/home/admin/bootstrap.sh"]
ENTRYPOINT ["./home/admin/bin/vegas"]