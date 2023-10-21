service-up:
	sudo docker-compose up --remove-orphans --build

service-down:
	sudo docker-compose down

.PHONY: service-up service-down
