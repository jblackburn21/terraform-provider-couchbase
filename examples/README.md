# Terraform Provider Couchbase Examples

## Run Locally

```bash
# RUn Couchbase locally
docker compose up -d

# Follow logs until server is running
docker compose logs -f
# Ctrl +c
```

## Examples

Clean when testing a new version of the provider

```bash
make clean
```

* [Bucket](./bucket/README.md)
