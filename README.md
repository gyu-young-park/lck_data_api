# Docker Makefile
Docker container is exposed on external port `8000`, so you can access on `http://localhost:8000/` 

1. docker image build
```sh
make build
```

2. docker container start
```sh
make start
```

3. remove container
```sh
make clean
```

4. stop container
```sh
make stop
```

5. restart container (after stop)
```
make restart
```

6. check
```
http://localhost:8000/v1/health/

Healthy check success!
```