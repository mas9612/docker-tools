# docker-tools

Before use, you should set `DOCKER_API_VERSION` environment variable.
```sh
$ export DOCKER_API_VERSION=`docker version | awk '/API version/ {print $3}' | head -1`
```

## image-remove
Delete specified images.

You can specify how many generations want to keep by `-generation` option.

```sh
$ image-remove -help
  -f    Force removal of the image
  -generation int
        Delete images older than this generation
  -name string
        Image name that you want to delete
$ image-remove -name alpine
```
