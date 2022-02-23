# M3O Proxy

The m3o proxy is a client for the M3O apis which you can run locally

## Usage

Populate the `M3O_API_TOKEN` environment variable

Install the m3o-proxy

```
go install github.com/m3o/m3o-proxy
```

Or download the latest [release](https://github.com/m3o/m3o-proxy/releases)

Run it like so

```
m3o-proxy
```

It will be running on `localhost:8080`. Now call helloworld.

```
curl http://localhost:8080/v1/helloworld/Call \
  -d '{"name": "Alice"}
```
