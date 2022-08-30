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

## Installation

Linux

```bash
curl -sfLo - https://github.com/justmiles/hole-punch/releases/download/v0.0.2/hole-punch_0.0.2_Linux_x86_64.tar.gz | tar -xzf - -C ~/.local/bin hp
```

Mac

```bash
curl -sfLo - https://github.com/justmiles/hole-punch/releases/download/v0.0.2/hole-punch_0.0.2_Darwin_x86_64.tar.gz | tar -xzf - -C ~/.local/bin hp
```

Mac (arm)

```bash
curl -sfLo - https://github.com/justmiles/hole-punch/releases/download/v0.0.2/hole-punch_0.0.2_Darwin_arm64.tar.gz | tar -xzf - -C ~/.local/bin hp
```

Windows

```cmd
# IDK, you're on your own here
```
