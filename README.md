# Tokenplay Service

This is the Tokenplay service

Generated with

```
micro new github.com/nicocesar/tokenplay --namespace=nicocesar --fqdn=tokenplay --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: tokenplay
- Type: srv
- Alias: tokenplay

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend etcd.

```
# install etcd
brew install etcd

# run etcd
etcd
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./tokenplay-srv
```

Build a docker image
```
make docker
```