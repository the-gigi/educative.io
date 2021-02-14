build:
	docker build . -t educative-go-course:latest

docker_tarball:
	tar -czvf tarball.tar.gz Dockerfile

data_tarball:
	tar -czpf data.tar.gz code

shell:
	docker run -it educative-go-course:latest bash
