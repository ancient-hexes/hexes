## Linting
```shell
docker run --rm -v $(pwd):/hexes -w /hexes/proto bufbuild/buf:1.6.0 lint
```

## Generating
```shell
rm -rf proto/go web/src/proto
docker run --rm -v $(pwd):/hexes -w /hexes/proto bufbuild/buf:1.6.0 generate
```
