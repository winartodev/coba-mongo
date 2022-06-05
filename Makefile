.PHONY: setup start_db stop_db remove run mod test

VOLUME_NAME = mongo_ecomerce_volume
CONFIG_NAME = mongo_ecomerce_config
CONTAINER_NAME = mongo_ecomerce
DATABASE_NAME = e_comerce

setup:
	docker volume create ${VOLUME_NAME}
	docker volume create ${CONFIG_NAME}
	docker run --name ${CONTAINER_NAME} \
	-v ${VOLUME_NAME}:/data/db \
	-v ${CONFIG_NAME}:/data/configdb \
	-p 27017:27017 \
	-d mongo:5.0 \

start_db:
	docker start ${CONTAINER_NAME}

stop_db:
	docker stop ${CONTAINER_NAME}

remove:
	docker container rm -f ${CONTAINER_NAME}
	docker volume rm ${VOLUME_NAME}
	docker volume rm ${CONFIG_NAME}

run: 
	go run ./...

test: 
	go test ./...

mod: 
	go mod tidy
	go mod download