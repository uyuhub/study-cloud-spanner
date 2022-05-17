# 説明

spanner 勉強用

# 立ち上げ方

#####　エミュレータ起動

```
docker-compose up -d
```

##### Spanner container 接続

```
make spanner
```

※TODO: なぜか database が存在しないとエラーになるので解消したい。  
とりあえず wrench create を先に実行すれば database が作成される。

# 困った時

##### 「OCI runtime create failed」「" 〜 permission denied": unknown.」 と言われたとき

<参考 URL>(https://www.y-hakopro.com/entry/2021/01/11/200130)

```
chmod +x spanner-init.sh.sh
```
