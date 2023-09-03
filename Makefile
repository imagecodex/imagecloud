build:
	go mod vendor
	docker build -t songjiayang/imagecloud:dev .
dev:
	docker run --rm --name imagecloud-dev -d -p 8080:8080 songjiayang/imagecloud:dev  
	open http://localhost:8080/example.jpg?x-amz-process=image/resize,w_800
stop:
	docker stop imagecloud-dev
test: 
	docker run --rm --name imagecloud-test -v $(CURDIR):/imagecloud -it songjiayang/govips:v0.2.0 sh -c "cd /imagecloud && go test ./..." 