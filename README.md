# CA-GO-Mission

## Run
```
# コンテナの作成
$ docker-compose up -d

# 起動したコンテナにログイン
$ docker exec -it db bash 

# MySQLを起動
$ mysql -u root -p  <pass>

# コンテナを修了
$ docker-compose down
```