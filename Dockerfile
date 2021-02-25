FROM golang
LABEL maintainer="stynxh@gmail.com"
# timezone 설정을 위함
ARG DEBIAN_FRONTEND=noninteractive
ENV TZ=Asia/Seoul

RUN apt update
RUN apt-get install -y tzdata

# golang web framework(gin) 설치
RUN go get -u github.com/gin-gonic/gin

# 웹 페이지 소스 파일 복사
RUN mkdir /go/src/webpage
COPY src /go/src/webpage

# WORK DIR 설정
WORKDIR /go/src/webpage

CMD ["go", "run", "main.go"]