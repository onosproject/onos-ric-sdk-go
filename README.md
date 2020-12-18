# onos-ric-sdk-go
[Go] Application SDK for ONOS RIC (µONOS Architecture)

The goal of this library is to make application development as easy as possible. To that end, the library should rely 
heavily on a set of newly established conventions that will result in certain default behaviours. 
To allow some applications to depart from these defaults, the library should be written in a modular 
fashion with high level abstractions and behaviours composed from lower-level ones. Most applications should be able
to rely on the top-level abstractions, but some apps may nned to instead utilize the lower-level abstraction.

## Usage

The SDK is managed using [Go modules]. To include the SDK in your Go application, add the `github.com/onosproject/onos-ric-sdk-go` module to your `go.mod`:

```
go get github.com/onosproject/onos-ric-sdk-go
```

## E2 API

The `github.com/onosproject/onos-ric-sdk-go/pkg/e2` provides clients and high-level interfaces for interacting
with E2 nodes. To create an E2 client:

```go
import "github.com/onosproject/onos-ric-sdk-go/pkg/e2"

...

config := e2.Config{
    AppID: "my-app",
    InstanceID: "my-app-1",
    SubscriptionService: e2.ServiceConfig{
        Host: "onos-e2sub",
    },
}

client, err := e2.NewClient(config)
if err != nil {
    ...
}
```

To subscribe to receive indications from an E2 node, first define the subscription:

```go
import "github.com/onosproject/onos-api/go/onos/e2sub/subscription"

...

var eventTrigger []byte // Encode the service model specific event trigger

details := subscription.SubscriptionDetails{
    E2NodeID: subscription.E2NodeID(nodeID),
    ServiceModel: subscription.ServiceModel{
        ID: subscription.ServiceModelID("test"),
    },
    EventTrigger: subscription.EventTrigger{
        Payload: subscription.Payload{
            Encoding: subscription.Encoding_ENCODING_PROTO,
            Data:     eventTrigger,
        },
    },
    Actions: []subscription.Action{
        {
            ID:   100,
            Type: subscription.ActionType_ACTION_TYPE_REPORT,
            SubsequentAction: &subscription.SubsequentAction{
                Type:       subscription.SubsequentActionType_SUBSEQUENT_ACTION_TYPE_CONTINUE,
                TimeToWait: subscription.TimeToWait_TIME_TO_WAIT_ZERO,
            },
        },
    },
}
```

Subscriptions support service model data definitions for event triggers and actions encoded in Protobuf or ASN.1.

Once the subscription is defined, use the `Subscribe` method to register the subscription and beging receiving indications. Indications are received in a Go channel:

```go
import "github.com/onosproject/onos-ric-sdk-go/pkg/e2/indication"

ch := make(chan indication.Indication)
_, err := client.Subscribe(context.TODO(), details, ch)

for ind := range ch {
    ...
}
```

The `Subscribe` method will return once the subscription has been registered with the subscription service. Once a subscription has been registered, it will not be unregistered until explicitly requested by the client. Changes to the state of the E2 node, subscription service, and termination points will be handled transparently by the core services and client implementation, which will work continuously to receive indications on the given channel.

A `context.Context` can be used to set a timeout for the subscription initialization:

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
defer cancel()
_, err := client.Subscribe(ctx, details, ch)
```

If the `Subscribe` call is successful, a `subcription.Context` will be returned. The context can be used to terminate the subscription by calling the `Close()` method:

```go
// Open the subscription
sub, err := client.Subscribe(context.Background(), details, ch)
...
// Close the subscription
err = sub.Close()
```

The `Close()` method will return once the subscription has been removed from the E2 Subscription service. The indications channel will be closed once the subscription has been closed, so it's safe to read from the indications channel in an indefinite loop.

## O1 API
...



## Watching Topology Changes
The µONOS RIC SDK provides a simple-to-use mechanism for application to initially scan, and then continuously monitor 
the topology for changes.

Application can do so by writing code similar to this:

```cgo
import (
    topoapi "github.com/onosproject/onos-api/go/onos/topo"
    "github.com/onosproject/onos-ric-sdk-go/pkg/topo"
)

    ...
    // Create a client entity that will be used to interact with the topology service.
    client, err := topo.NewClient(&topo.ServiceConfig{...})

    // Allocate channel on which topology events will be received
    ch := make(chan topoapi.Event)
    
    // Start watching for topology changes that match the following filters:
    //  - entities only (not relations)
    //  - entities of kind "eNB" only
    //  - existing entities or entities being added
    err = client.Watch(context.Background(), ch,
            WithTypeFilter(topoapi.Object_ENTITY), WithKindFilter("eNB"),
            WithEventFilter(topoapi.EventType_NONE, topoapi.EventType_ADDED))
    ...
    
    // Read from the channel; likely in a go routine
    for event := range ch {
        // Process each event
        ...
    }
    
    ...
    
    // Close the channel when stopping
    close (ch)
		
```

To narrow down the type of events to be received, the application can specify the following filter options to the `Watch` call:

* `WithTypeFilter` - list of object types, `ENTITY`, `RELATION`, `KIND`
* `WithKindFilter` - list of specific entity or relationship kinds, e.g. `gNB`, `switch`, `controller`
* `WithEventFilter` - list of event types `NONE` (pre-existing item), or newly `ADDED`, `UPDATED`, `REMOVED`

[Go]: https://golang.org/
[Go modules]: https://golang.org/ref/mod
