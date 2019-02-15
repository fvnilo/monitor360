build: Dockerfile
	docker build -t nyfanilo/monitor360:latest .

deploy: build
	docker push nyfanilo/monitor360:latest && \
	ssh -t amazon './start-monitor.sh'
	
run-local: build
	docker run -p 80:80  nyfanilo/monitor360:latest

dev: cmd/server/main.go
	go run cmd/server/main.go