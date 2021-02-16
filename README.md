# CA-GO-Mission

## Run
```
# コンテナの作成
$ docker-compose up -d

# 起動したコンテナにログイン
$ docker exec -it Docker-mysql_mysql_1 bash -p

# MySQLを起動
$ mysql -u root -p  <pass>

# コンテナを修了
$ docker-compose down
```