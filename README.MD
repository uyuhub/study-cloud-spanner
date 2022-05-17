# 概要

spanner 勉強用です。
エミュレーターが起動できます。

# 使い方

##### 環境変数の設定(direnv 利用)

```
make env
```

##### エミュレータ起動

```
docker-compose up -d
```

##### Spanner container 接続

```
make spanner
```

##### ddl 実行(初回のみ)

```
make createdb
```

##### db schema 再構築

```
make resetdb
```

# 困った時

##### 「OCI runtime create failed」「" 〜 permission denied": unknown.」 と言われたとき

<参考 URL>(https://www.y-hakopro.com/entry/2021/01/11/200130)

```
chmod +x spanner-init.sh
```
