build:
	go mod vendor
	docker build -t imagecodex/imagecloud:dev .
build-vips:
	docker build -f Dockerfile.vips -t imagecodex/govips:v0.3.0 .
dev:
	docker run --rm --name imagecloud-dev -d -p 8080:8080 -p 8081:8080 imagecodex/imagecloud:dev  
	open http://localhost:8080/example.jpg?x-amz-process=image/resize,w_800
stop:
	docker stop imagecloud-dev
test: 
	go mod vendor
	docker run --rm --name imagecloud-test -v $(CURDIR):/imagecloud -i imagecodex/govips:v0.3.0 sh -c "cd /imagecloud && go test ./..." 
lint:
	go mod vendor
	docker run --rm --name imagecloud-test -v $(CURDIR):/imagecloud -i imagecodex/govips:v0.3.0 sh -c "cd /imagecloud && golangci-lint run ./..." 