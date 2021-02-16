# CA-GO-Mission


docker run --name db -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root mysql

<!-- ## Start database
```
sudo docker run -d --name db -e MYSQL_ROOT_PASSWORD=mysql -e MYSQL_DATABASE=DB -v ./user.sql:/docker-entrypoint-initdb.d/user.sql mysql
``` -->

<!-- ## Start mysql
```
# イメージのビルド
$ Docker-Compose build

# コンテナの作成
$ Docker-Compose up -d

# 起動したコンテナにログイン
$ docker exec -it Docker-mysql_mysql_1 bash -p

# MySQLを起動
$ mysql -u root -p -h 127.0.0.1
``` -->