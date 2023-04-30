# Skadi

Low-level socket communication using Redis

## How to run

### Docker

The command to start a server in Docker:

```shell
make up
```

### Local

The command to start the server:

```shell
make server-run
```

The command to start the client:

```shell
make client-run
```

## Features

- [X] Custom binary communication protocol
- [X] Server and client on websockets
- [X] Redis for storage
- [ ] Server shutdown on client side
- [ ] The server creates a .pid file which contains its PID. Deleting the file will force the server to close.
- [ ] Only one server instance can run at a time (Not including Docker)
