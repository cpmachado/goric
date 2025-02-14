# goric

> "Weddings are basically funerals with cake."
>
> by Rick, from "Rick and Morty"

goric is an implementation of [ric](https://github.com/cpmachado/ric) in go

## Status

Initial development

Implements:
- Task 1: Get Host name
- Task 2: Get IP from a given host
- Task 3, 4, 5: Simple UDP client that sends "Hello!\n", and expects reply

## Usage

```text
goric is a WIP clone of 'ric'
Usage: goric [OPTIONS]
  -d string
    	HOSTNAME (default "localhost")
  -n	Task 1: hostname
  -p int
    	PORT (default 1337)
  -u	Task 3,4,5: udp client to contact UDP echo server
  -v	Display version and exit
  -w	Task 2: nslook
```


## Versioning

Before v1.0.0:
- Each minor version is a step in the original guide.
