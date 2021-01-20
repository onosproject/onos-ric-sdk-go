// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: LicenseRef-ONF-Member-1.0

package store

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/onosproject/onos-ric-sdk-go/pkg/config/event"

	"github.com/stretchr/testify/assert"
)

var data = `{
  "interfaces": {
    "interface": [
      {
        "name": "admin",
        "config": {
          "name": "admin"
        }
      }
    ]
  },
  "system": {
    "aaa": {
      "authentication": {
        "admin-user": {
          "config": {
            "admin-password": "password"
          }
        },
        "config": {
          "authentication-method": [
            "openconfig-aaa-types:LOCAL"
          ]
        }
      }
    },
    "clock": {
      "config": {
        "timezone-name": "Europe/Dublin"
      }
    },
    "config": {
      "hostname": "replace-device-name",
      "domain-name": "opennetworking.org",
      "login-banner": "This device is for authorized use only",
      "motd-banner": "replace-motd-banner"
    },
    "state" : {
      "boot-time": "1575415411",
      "current-datetime": "2019-12-04T10:00:00Z-05:00",
      "hostname": "replace-device-name",
      "domain-name": "opennetworking.org",
      "login-banner": "This device is for authorized use only",
      "motd-banner": "replace-motd-banner"

    },
    "openflow": {
      "agent": {
        "config": {
          "backoff-interval": 5,
          "datapath-id": "00:16:3e:00:00:00:00:00",
          "failure-mode": "SECURE",
          "inactivity-probe": 10,
          "max-backoff": 10
        }
      },
      "controllers": {
        "controller": [
          {
            "config": {
              "name": "main"
            },
            "connections": {
              "connection": [
                {
                  "aux-id": 0,
                  "config": {
                    "address": "192.0.2.10",
                    "aux-id": 0,
                    "port": 6633,
                    "priority": 1,
                    "source-interface": "admin",
                    "transport": "TLS"
                  },
                  "state": {
                    "address": "192.0.2.10",
                    "aux-id": 0,
                    "port": 6633,
                    "priority": 1,
                    "source-interface": "admin",
                    "transport": "TLS"
                  }
                },
                {
                  "aux-id": 1,
                  "config": {
                    "address": "192.0.2.11",
                    "aux-id": 1,
                    "port": 6653,
                    "priority": 2,
                    "source-interface": "admin",
                    "transport": "TLS"
                  },
                  "state": {
                    "address": "192.0.2.11",
                    "aux-id": 1,
                    "port": 6653,
                    "priority": 2,
                    "source-interface": "admin",
                    "transport": "TLS"
                  }
                }

              ]

            },
            "name": "main"
          },
          {
            "config": {
              "name": "second"
            },
            "connections": {
              "connection": [
                {
                  "aux-id": 0,
                  "config": {
                    "address": "192.0.3.10",
                    "aux-id": 0,
                    "port": 6633,
                    "priority": 1,
                    "source-interface": "admin",
                    "transport": "TLS"
                  },
                  "state": {
                    "address": "192.0.3.10",
                    "aux-id": 0,
                    "port": 6633,
                    "priority": 1,
                    "source-interface": "admin",
                    "transport": "TLS"
                  }
                },
                {
                  "aux-id": 1,
                  "config": {
                    "address": "192.0.3.11",
                    "aux-id": 1,
                    "port": 6653,
                    "priority": 2,
                    "source-interface": "admin",
                    "transport": "TLS"
                  },
                  "state": {
                    "address": "192.0.3.11",
                    "aux-id": 1,
                    "port": 6653,
                    "priority": 2,
                    "source-interface": "admin",
                    "transport": "TLS"
                  }
                }

              ]

            },
            "name": "second"
          }
        ]
      }
    }
  }
}
`

const (
	Key1 = "/system/config/hostname"
	Val1 = "test-hostname"

	Key2 = "/interfaces/interface[name=admin]/config/name"
	Val2 = "admin"

	Key3 = "/system/openflow/controllers/controller[name=second]/connections/connection[aux-id=0]/config/port"
	Val3 = 6634
)

func TestStore(t *testing.T) {
	dataByte := []byte(data)
	store := NewConfigStore()
	ch := make(chan event.Event)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err := json.Unmarshal(dataByte, &store.ConfigTree)
	assert.NoError(t, err)
	err = store.Watch(ctx, ch)
	assert.NoError(t, err)
	err = store.Put(Key1, Entry{
		Key:   Key1,
		Value: Val1,
	})
	assert.NoError(t, err)

	event := <-ch
	assert.Equal(t, event.Key, Key1)
	assert.Equal(t, event.Value, Val1)

	val, err := store.Get(Key1)
	assert.NoError(t, err)
	assert.Equal(t, val.Value, Val1)
	assert.NoError(t, err)

	val, err = store.Get(Key2)
	assert.NoError(t, err)
	assert.Equal(t, val.Value, Val2)

	entry, err := store.Get(Key3)
	assert.NoError(t, err)
	assert.Equal(t, entry.Value, float64(6633))
	err = store.Put(Key3, Entry{Key: Key3, Value: float64(Val3)})
	assert.NoError(t, err)
	event = <-ch
	assert.Equal(t, event.Key, Key3)
	assert.Equal(t, event.Value, float64(Val3))
	val, err = store.Get(Key3)
	assert.NoError(t, err)
	assert.Equal(t, val.Value, float64(6634))

}
