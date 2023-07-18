build:
	docker build -t ip-info-app .

run:
	docker run --env-file .env -p $(port) ip-info-app
