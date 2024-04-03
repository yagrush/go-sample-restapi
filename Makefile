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
	docker container exec -it app sh
