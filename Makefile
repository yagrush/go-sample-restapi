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

# downと同時に名前付きボリュームを全て削除する
down-v:
	docker-compose down --volume

# 未使用のDcokerオブジェクトを掃除する
docker-prune:
	docker system prune -a