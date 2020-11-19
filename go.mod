module github.com/onosproject/onos-ric-sdk-go

go 1.15

require (
	github.com/gogo/protobuf v1.3.1
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/google/go-cmp v0.5.1 // indirect
	github.com/google/uuid v1.1.2
	github.com/onosproject/config-models/modelplugin/ric-1.0.0 v0.0.0-20201118220729-1f227212be88
	github.com/onosproject/gnxi-simulators v0.6.4 // indirect
	github.com/onosproject/onos-e2sub v0.6.4-0.20201112225438-8ff953e41a09
	github.com/onosproject/onos-e2t v0.6.7-0.20201112232226-f90757e4b4c0
	github.com/onosproject/onos-lib-go v0.6.25
	github.com/openconfig/gnmi v0.0.0-20200617225440-d2b4e6a45802
	github.com/openconfig/goyang v0.2.1
	github.com/openconfig/ygot v0.8.11
	golang.org/x/net v0.0.0-20200707034311-ab3426394381
	golang.org/x/sys v0.0.0-20200803210538-64077c9b5642 // indirect
	golang.org/x/text v0.3.3 // indirect
	google.golang.org/genproto v0.0.0-20200731012542-8145dea6a485
	google.golang.org/grpc v1.33.2
	gopkg.in/yaml.v2 v2.3.0 // indirect
)

replace github.com/onosproject/config-models/modelplugin/ric-1.0.0 => /Users/adibrastegarnia/go/src/github.com/onosproject/config-models/modelplugin/ric-1.0.0

replace github.com/onosproject/config-models => /Users/adibrastegarnia/go/src/github.com/onosproject/config-models
