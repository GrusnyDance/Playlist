# на моем линуксе докер запускается c docker compose, без dash
run:
	docker-compose up -d --build --force-recreate

run_not_detached:
	docker compose up --build --force-recreate

stop:
	docker-compose down

clean:
# возможно на маке не нужны отражающие символы
	docker stop $$(docker ps -a -q) || true
	docker rm $$(docker ps -a -q) || true
	docker rmi $$(docker images -a -q) || true

docker-test-client:
	docker build --no-cache -t client -f DockerfileClient .
	docker run --privileged client

docker-test-server:
	docker build --no-cache -t server -f DockerfileServer .
	docker run --privileged server
