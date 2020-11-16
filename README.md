# onos-ric-sdk-go
Golang Application SDK for ONOS RIC (µONOS Architecture)

The goal of this library is to make application development as easy as possible. To that end, the library should rely 
heavily on a set of newly established conventions that will result in certain default behaviours. 
To allow some applications to depart from these defaults, the library should be written in a modular 
fashion with high level abstractions and behaviours composed from lower-level ones. Most applications should be able
to rely on the top-level abstractions, but some apps may nned to instead utilize the lower-level abstraction.

The library mey need to track its internal state and for this purpose there will be an entity called `ric.ApplicationContext`
that will be established via a call to:

`ric.Begin(options ric.Options) ric.ApplicationContext`

## E2 API
To interact with E2 Nodes application must subscribe to a set of E2 SM messages via at least one, but possibly several calls to:

`e2.Subscribe(subscription e2.Subscription, ch chan<- indication.Indication) error`

In response to this call, the library will issue a subscription request to the E2 subscription manager and then will 
start to internally manage a set of connections to the available E2 termination nodes by listening to notifications from 
the E2 subscription manager to tear-down and setup connections as necessary. It will also begin to track which connections should
be used for which E2 nodes.

Incoming messages resulting from this subscription will be routed to the specified channel. Different subscription requests can indicate different channels.

To send a message to an E2 node, the application can call the following:

`e2.Send(nodeId e2.NodeId, message e2.Message) error`

## O1 API
...
