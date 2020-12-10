// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package topo

import (
	"github.com/golang/mock/gomock"
	topoapi "github.com/onosproject/onos-api/go/onos/topo"
	"github.com/onosproject/onos-lib-go/pkg/errors"
	"github.com/onosproject/onos-ric-sdk-go/tests/mocks"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestTopoWatch(t *testing.T) {
	ctrl := gomock.NewController(t)

	e1 := &topoapi.Object_Entity{Entity: &topoapi.Entity{KindID: "eNB"}}
	e2 := &topoapi.Object_Entity{Entity: &topoapi.Entity{KindID: "gNB"}}
	er := &topoapi.Object_Relation{Relation: &topoapi.Relation{KindID: "foo"}}

	o1 := topoapi.Object{ID: "1", Type: topoapi.Object_ENTITY, Obj: e1}
	o2 := topoapi.Object{ID: "2", Type: topoapi.Object_ENTITY, Obj: e1}
	o3 := topoapi.Object{ID: "3", Type: topoapi.Object_ENTITY, Obj: e2}
	r1 := topoapi.Object{ID: "r", Type: topoapi.Object_RELATION, Obj: er}

	// Prepare mock watch stream
	stream := mocks.NewMockTopo_WatchClient(ctrl)
	stream.EXPECT().Recv().Return(&topoapi.WatchResponse{Event: topoapi.Event{Object: o1, Type: topoapi.EventType_NONE}}, nil)
	stream.EXPECT().Recv().Return(&topoapi.WatchResponse{Event: topoapi.Event{Object: o2, Type: topoapi.EventType_ADDED}}, nil)
	stream.EXPECT().Recv().Return(&topoapi.WatchResponse{Event: topoapi.Event{Object: o3, Type: topoapi.EventType_ADDED}}, nil)
	stream.EXPECT().Recv().Return(&topoapi.WatchResponse{Event: topoapi.Event{Object: r1, Type: topoapi.EventType_ADDED}}, nil)
	stream.EXPECT().Recv().Return(&topoapi.WatchResponse{Event: topoapi.Event{Object: o1, Type: topoapi.EventType_REMOVED}}, nil)
	stream.EXPECT().Recv().Return(nil, io.EOF)

	// Start mock topology service
	mockClient := mocks.NewMockTopoClient(ctrl)
	mockClient.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(stream, nil)

	client := testClient(mockClient)

	ch := make(chan topoapi.Event)
	err := client.Watch(ch,
		TypeFilter(topoapi.Object_ENTITY), KindFilter("eNB"),
		EventFilter(topoapi.EventType_NONE, topoapi.EventType_ADDED))
	assert.NoError(t, err, "unable to start topology watch")

	e := <- ch
	assert.Equal(t, topoapi.ID("1"), e.Object.ID)
	e = <- ch
	assert.Equal(t, topoapi.ID("2"), e.Object.ID)
	e = <- ch
	assert.Equal(t, topoapi.Object_UNSPECIFIED, e.Object.Type)
}

func TestTopoWatchFail(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Start mock topology service
	mockClient := mocks.NewMockTopoClient(ctrl)
	mockClient.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(nil, errors.NewTimeout("zzzz"))

	client := testClient(mockClient)
	ch := make(chan topoapi.Event)
	err := client.Watch(ch, TypeFilter(topoapi.Object_ENTITY), KindFilter("eNB"))
	assert.Error(t, err)
}

func TestTopoWatchStreamFail(t *testing.T) {
	ctrl := gomock.NewController(t)

	e1 := &topoapi.Object_Entity{Entity: &topoapi.Entity{KindID: "eNB"}}
	o1 := topoapi.Object{ID: "1", Type: topoapi.Object_ENTITY, Obj: e1}

	// Prepare mock watch stream
	stream := mocks.NewMockTopo_WatchClient(ctrl)
	stream.EXPECT().Recv().Return(&topoapi.WatchResponse{Event: topoapi.Event{Object: o1, Type: topoapi.EventType_NONE}}, nil)
	stream.EXPECT().Recv().Return(nil, errors.NewTimeout("zzzz"))
	stream.EXPECT().Recv().Return(nil, io.EOF)

	// Start mock topology service
	mockClient := mocks.NewMockTopoClient(ctrl)
	mockClient.EXPECT().Watch(gomock.Any(), gomock.Any()).Return(stream, nil)

	client := testClient(mockClient)

	ch := make(chan topoapi.Event)
	err := client.Watch(ch, TypeFilter(topoapi.Object_ENTITY), KindFilter("eNB"))
	assert.NoError(t, err, "unable to start topology watch")

	e := <- ch
	assert.Equal(t, topoapi.ID("1"), e.Object.ID)
	e = <- ch
	assert.Equal(t, topoapi.Object_UNSPECIFIED, e.Object.Type)
}


// TODO: add negative tests


