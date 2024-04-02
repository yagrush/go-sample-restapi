# go-sample-restapi

自分流のクリーンアーキテクチャを育てていくリポジトリ

---

## run
```
make run
```

### access to sample API: calcAddInt64
http://localhost:8089/sample/calcAddInt64?a=5&b=9

----

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
