docker build -t entrypoint .
docker run -it --rm -v $PWD:/workspace entrypoint