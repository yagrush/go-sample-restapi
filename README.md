# go-sample-restapi

自分流のクリーンアーキテクチャを育てていくリポジトリ

---

細かいhow to useは `Makefile` をご参照ください

## build and run
```
make start
```

## access to sample API: calcAddInt64
http://localhost:8083/api/v1/sample/calcAddInt64?a=1&b=237


## tail DB logs
```
make tail-db
```

## tail KVS logs
```
make tail-kvs
```

## destroy
```
make down
```

## destroy with DB reest
```
make down-v
```
KVSのデータリセットは `kvs/data` を削除してください

---

## test app
```
make test-app
```

---

## memo

### キャッシュを無視してgo install
```
GOPROXY=direct go install <package>@latest
```
