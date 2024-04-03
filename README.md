# go-sample-restapi

自分流のクリーンアーキテクチャを育てていくリポジトリ

---

## build and run
```
make start
```

## access to sample API: calcAddInt64
http://localhost:8083/api/v1/sample/calcAddInt64?a=1&b=237

## destroy
```
make down
```

---

## test
```
make test
```

---

## memo

### キャッシュを無視してgo install
```
GOPROXY=direct go install <package>@latest
```
