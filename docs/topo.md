## Topo API
The `github.com/onosproject/onos-ric-sdk-go/pkg/topo` provides a high level interface for applications to interact
with [onos-topo] subsystem. The current interface allows applications to *Create*, *Update*, *Get*, and *List*  topology entities and their relations and *Watch* topology changes.


### Create A Topo Client
The following code snippet shows how to create an instance of topo SDK client:

```go
import (
   toposdk "github.com/onosproject/onos-ric-sdk-go/pkg/topo"
   topoapi "github.com/onosproject/onos-api/go/onos/topo"
)

topoClient, err := toposdk.NewClient()
if err != nil {
   return err
}

```
### Create or Update Topo Objects
The current SDK provides *Create* and *Update* methods which allows
creating new topology objects or updating existing ones. As an example, the following
code snippet shows how to create an E2 node entity object using *Create* method:

```go
 ...

object: = &topoapi.Object {
         ID:   e2NodeID,
         Type: topoapi.Object_ENTITY,
         Obj: & topoapi.Object_Entity {
         	Entity: & topoapi.Entity {
         KindID: topoapi.ID(topoapi.E2NODE),
         },
     },
        Aspects: make(map[string]*gogotypes.Any),
        Labels: map[string] string {},
}

// Define E2 node aspects
e2NodeAspects := &topoapi.E2Node{
	Servicemodels: <Service models Information>    // Service models information
}
err := object.SetAspect(e2NodeAspects)
err = topoClient.Create(context.TODO(), object)

```

To update an existing Object, *Get* method can be used to retrieve
an existing object and then after updating its fields, the modified topo Object
should be passed to *Update* method.


### Get or List Topo Objects
To get a topo object, the topology object ID should be passed to Get method as follows:

```go
topoObject, err := topoClient.Get(context.TODO(), topoObjectID)
```

*List* method can be used to obtain a collection of objects. *List* method provides [Topology filters]
option to allow listing specific entities or relations. For example, the following code snippet lists all of *Control* relations in topo

```go
controlRelationFilter := &topoapi.Filters{
         KindFilter: &topoapi.Filter{
             Filter: &topoapi.Filter_Equal_{
               Equal_: &topoapi.EqualFilter{
                 Value: topoapi.CONTROLS,
         },
       },
     },
  }
objects, err := topoClient.List(context.TODO(), controlRelationFilter)  

```



### Watch Topo Changes
Current SDK provides a *Watch* method that allows applications to
monitor topology changes using Go channels. *Watch* method provides [Topology filters] option that allows  monitoring 
changes for specific entities or relations. For example, the following code snippet shows 
how to use *Watch* method to monitor *Control* relations

```go

controlRelationFilter := &topoapi.Filters{
		KindFilter: &topoapi.Filter{
			Filter: &topoapi.Filter_Equal_{
				Equal_: &topoapi.EqualFilter{
					Value: topoapi.CONTROLS,
				},
			},
		},
	}

ch := make(chan topoapi.Event)
err := topoClient.Watch(context.TODO(), ch, toposdk.WithWatchFilters(controlRelationFilter))

for event := range ch {
   // Process each event
}

// Close the channel when stopping
close (ch)

```

[onos-topo]: https://github.com/onosproject/onos-topo
[Topology filters]: https://github.com/onosproject/onos-topo#filters