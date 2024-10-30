## test

docker build -f ./test/Dockerfile -t 117503445/dev-test .
docker run -v $PWD/test:/workspace --rm -it 117503445/dev-test
docker run --privileged -v $PWD/test:/workspace --rm -it 117503445/dev-test
