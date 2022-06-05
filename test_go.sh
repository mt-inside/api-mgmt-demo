pushd foo
just build
just package
docker run -ti -p8090:8090 docker.io/mtinside/foo:0.1
