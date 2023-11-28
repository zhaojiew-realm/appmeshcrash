build-release:
	docker build -t zhaojiew/appmeshload:latest .
	docker push zhaojiew/appmeshload:latest

build-debug:
	docker build -f dockerfile_debug -t zhaojiew/appmeshload:debug .
	docker push zhaojiew/appmeshload:debug