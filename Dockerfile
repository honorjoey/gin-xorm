FROM golang:latest

LABEL maintainer="jamesqiao293@gmail.com"
LABEL description="a example of gin and xorm"
ENV GOPROXY https://goproxy.cn,direct
ENV PROJECT_PATH $GOPATH/src/github.com/honorjoey/gin-xorm
WORKDIR $PROJECT_PATH
COPY . $PROJECT_PATH
RUN go get -u github.com/swaggo/swag/cmd/swag@v1.6.5 && mv $GOPATH/bin/swag /usr/local/go/bin \
    && swag init && go build .

EXPOSE 8080

ENTRYPOINT ["./gin-xorm"]
CMD []