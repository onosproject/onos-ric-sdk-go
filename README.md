# onos-ric-sdk-go
[Go] Application SDK for ONOS RIC (µONOS Architecture)

The goal of this library is to make application development as easy as possible. To that end, the library should rely 
heavily on a set of newly established conventions that will result in certain default behaviours. 
To allow some applications to depart from these defaults, the library should be written in a modular 
fashion with high level abstractions and behaviours composed of lower-level ones. Most applications should be able
to rely on the top-level abstractions, but some apps may need to instead utilize the lower-level abstraction.

## Installation

The SDK is managed using [Go modules]. To include the SDK in your Go application, add the `github.com/onosproject/onos-ric-sdk-go` module to your `go.mod`:

```
go get github.com/onosproject/onos-ric-sdk-go
```

## Usage

For the detail usage of each API, you can refer to the following links:

[E2 API Usage](docs/e2.md)

[O1 API Usage](docs/o1.md)

[Topo API Usage](docs/topo.md)

[A1 API Usage](docs/a1.md)

[Go]: https://golang.org/
[Go modules]: https://golang.org/ref/mod
[onos-config]: https://github.com/onosproject/onos-config
