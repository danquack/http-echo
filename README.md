# http-echo
Simple server that logs body to console

### Execute
```sh
docker run --rm -p 8090:8090 ghcr.io/danquack/http-echo
curl http://localhost:8090 -d 'testing'
```

## Environment variable
- LOG_HEADERS: (default: true) `true/false` to log the headers that are posted.
