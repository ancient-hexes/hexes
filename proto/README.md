## Linting
```shell
docker run --rm -v $(pwd):/proto -w /proto bufbuild/buf:1.6.0 lint
```

## Generating
```shell
rm -rf generated
docker run --rm -v $(pwd):/proto -w /proto bufbuild/buf:1.6.0 generate
```
