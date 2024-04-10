generate:
	git ls-files '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && make generate'

tidy:
	git ls-files '**/*go.mod' -z | xargs -0 -I{} bash -xc 'cd $$(dirname {}) && go mod tidy'

test-app:
	go test -race -shuffle=on ./app/...

up:
	docker-compose up

start:
	docker-compose up --build -d

stop:
	docker-compose stop

down:
	docker-compose down

log:
	docker-compose logs -f

sh-app:
	docker container exec -it sample-app sh

sh-db:
	docker container exec -it sample-db bash

sh-kvs:
	docker container exec -it sample-kvs bash

tail-db:
	docker container exec -it sample-db tail -f /tmp/log/db_query.log /tmp/log/mysqld.log

tail-kvs:
	docker container exec -it sample-kvs tail -f /var/log/redis/redis.log

# downと同時に名前付きボリュームを全て削除する
down-v:
	docker-compose down --volume

# 未使用のDcokerオブジェクトを掃除する
docker-prune:
	docker system prune -a