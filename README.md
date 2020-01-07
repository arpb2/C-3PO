# C-3PO

Backend for the ARPB2 project


## Development

Clone the repository at `$GOPATH/go/src/github.com/arpb2` and start developing!

## Running

Simply run (standing at the root of the project)

```bash
docker-compose up # Use --build if you need to rebuild the image
```

After that, simply access them through the exposed ports in the compose

```bash
# API
curl localhost:5555/ping
```

Or as websites:
- [Grafana](http://localhost)
- [CAdvisor](http://localhost:8080)
