# go-opensea

<img src="logo.jpg" alt="OpenSea Go SDK Logo" width="200"/>

I started this Go opensea SDK because the official openapi spec on the Opensea website is incorrect.

This is not exhaustive—I'll be adding methods as I need them

## Generating the SDK

The project uses `oapi-codegen` to generate the SDK from the open api spec.

To generate the client run

```shell
make generate
```

See `examples` for usage.

## Docs

https://stein-f.github.io/go-opensea