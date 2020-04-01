# Tokenplay Service

This is the Tokenplay service

Generated with

```
micro new github.com/nicocesar/tokenplay --namespace=nicocesar --fqdn=tokenplay --type=service
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: tokenplay
- Type: service
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
./tokenplay-service
```

Build a docker image
```
make docker
```