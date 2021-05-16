all: docker
ui: clean
	npm run build --prefix ./ui/
server: ui
	mkdir docker/dist
	env GOOS=linux CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w -extldflags "-static"' -o docker/dist/boardgamesurvey ./main.go
dist: server
	cp -R static docker/dist/
docker: dist
	docker build -t boardgamesurvey:v0.1 docker/
clean: 
	rm -rf docker/dist
	rm -rf bin