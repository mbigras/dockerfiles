# dockerfiles

> Examples of Dockerfiles for different languages

## Usage

```
git clone https://github.com/mbigras/dockerfiles
cd dockerfiles/go
docker build -t fooimage .
docker run \
	--name=foocontainer \
	--detach \
	--publish=5000:5000 \
	--env=PORT=5000 \
	fooimage
curl localhost:5000/hello/world
docker rm -f foocontainer
```
