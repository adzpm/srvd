# srvd

**Web-server for fast serving directory on `host`:`port`**

## Flags

| Flag | Description        | Default     |
|------|--------------------|-------------|
| `-d` | Directory to serve | `./`        |
| `-h` | Host to listen     | `127.0.0.1` |
| `-p` | Port to listen     | `8080`      |

## Run from source

```bash
go run github.com/adzpm/srvd@latest -d /path/to/directory -p 8085
```