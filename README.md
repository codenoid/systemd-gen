# systemd-gen

## Download

1. Download binary (Built on Ubuntu 18.04) from [release page](https://github.com/codenoid/systemd-gen/releases/download/0.0.1/systemd-gen)
2. `go get -u github.com/codenoid/systemd-gen`

## How to use it

```bash
# username is logged username
$ sudo systemd-gen service-name username '/path/to/project/dir' '/usr/bin/python3 something.py'
Service File Created.
```