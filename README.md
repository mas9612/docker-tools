# docker-tools

## image-remove
Delete specified images.

You can specify how many generations want to keep by `-generation` option.

```
$ image-remove -help
  -f    Force removal of the image
  -generation int
        Delete images older than this generation
  -name string
        Image name that you want to delete
$ image-remove -name alpine
```
