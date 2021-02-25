# Docker build 정보
- 테스트   
docker build -t goweb .
docker run -it -p 9999:80 --name goweb goweb   

- 운영   
docker build -t goweb .
docker run -d -p 9999:80 --name goweb goweb       

# 참고
- gin 웹프레임워크 : https://gin-gonic.com/ko-kr/docs/