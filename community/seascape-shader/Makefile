# go get -u github.com/gopxl/pixel-examples
# cd ~/go/src/github.com/gopxl/pixel-examples/community/seascape-shader

all:
	go build
	./seascape-shader

push:
	git pull
	git add --all
	-git commit -a -s
	git push

update:
	git pull

diff:
	git diff
