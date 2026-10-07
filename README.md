# tinyredis

A tiny Redis-like key-value store written from scratch in Go — no external libraries.

It runs a TCP server where each client connection gets its own goroutine,
and speaks a simple line-based protocol.

## Status

- [x] Echo server (TCP, goroutine per connection)
- [ ] Command protocol: `GET` / `SET` / `DEL`
- [ ] Concurrent in-memory store
- [ ] Graceful shutdown
- [ ] Persistence (append-only file)

## Run

```bash
go run ./cmd/tinyredis
```

Then connect with any TCP client and type — everything you send comes back:

```bash
$ nc localhost 7379
hello
hello
```

## Layout

```
cmd/tinyredis/     # main: starts the server
internal/server/   # TCP accept loop + per-connection handling
```