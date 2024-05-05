build:
	go build src/*.go

run: build
	./src/main

test:
	go test src/*.go

docker:
	docker build -t {{ cookiecutter.project_slug }} .

docker-run:
	sudo docker run -p 8080:8080 -it {{ cookiecutter.project_slug }}
