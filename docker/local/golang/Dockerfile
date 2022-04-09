FROM golang:1.15.2

COPY . /go/src

WORKDIR /go/src
ADD . /go/src/

# go get時にtest実行されgccライブライ関連でエラーになるのでテスト実行させない
RUN export CGO_ENABLED=0

RUN go get github.com/gin-gonic/gin@v1.7.4 \
	&& go get gorm.io/gorm@v1.21.16 \
	&& go get gorm.io/driver/mysql \
	&& go get github.com/joho/godotenv@v1.4.0 \
	&& go get github.com/fgrosse/goldi/goldigen

EXPOSE 3000