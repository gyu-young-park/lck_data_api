LCK_DATE_API_VER=v1
LCK_DATE_API_NAME=lck-data-api
LCK_DATE_API_IMAGE_TAG="$(LCK_DATE_API_NAME):$(LCK_DATE_API_VER)"
EXTERNAL_PORT=8000
CONTAINER_PORT=8080

build:
	docker build -t $(LCK_DATE_API_NAME) --tag $(LCK_DATE_API_IMAGE_TAG) .

start:
	docker run -d --name $(LCK_DATE_API_NAME) -p $(EXTERNAL_PORT):$(CONTAINER_PORT) $(LCK_DATE_API_IMAGE_TAG)

clean:
	make stop
	docker rm $(LCK_DATE_API_NAME)

stop:
	docker stop $(LCK_DATE_API_NAME)

restart:
	docker restart $(LCK_DATE_API_NAME)