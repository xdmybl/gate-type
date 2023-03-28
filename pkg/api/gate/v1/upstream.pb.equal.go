// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/xdmybl/gate-type/proto/gate/v1/upstream.proto

package v1

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *UpstreamSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*UpstreamSpec)
	if !ok {
		that2, ok := that.(UpstreamSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if h, ok := interface{}(m.GetCommonInfo()).(equality.Equalizer); ok {
		if !h.Equal(target.GetCommonInfo()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetCommonInfo(), target.GetCommonInfo()) {
			return false
		}
	}

	if m.GetLbAlg() != target.GetLbAlg() {
		return false
	}

	if h, ok := interface{}(m.GetSslConfigurations()).(equality.Equalizer); ok {
		if !h.Equal(target.GetSslConfigurations()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetSslConfigurations(), target.GetSslConfigurations()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetConnPoll()).(equality.Equalizer); ok {
		if !h.Equal(target.GetConnPoll()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetConnPoll(), target.GetConnPoll()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHealthCheckRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHealthCheckRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHealthCheckRef(), target.GetHealthCheckRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetStatefulSession()).(equality.Equalizer); ok {
		if !h.Equal(target.GetStatefulSession()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetStatefulSession(), target.GetStatefulSession()) {
			return false
		}
	}

	if len(m.GetEndpoints()) != len(target.GetEndpoints()) {
		return false
	}
	for idx, v := range m.GetEndpoints() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetEndpoints()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetEndpoints()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *ConnPoll) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*ConnPoll)
	if !ok {
		that2, ok := that.(ConnPoll)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetMaxRequestsPerConnection() != target.GetMaxRequestsPerConnection() {
		return false
	}

	if m.GetMaxConnections() != target.GetMaxConnections() {
		return false
	}

	if m.GetMaxRequests() != target.GetMaxRequests() {
		return false
	}

	if m.GetMaxPendingRequests() != target.GetMaxPendingRequests() {
		return false
	}

	if strings.Compare(m.GetOutboundSourceAddress(), target.GetOutboundSourceAddress()) != 0 {
		return false
	}

	return true
}

// Equal function
func (m *WeightedUpstreamList) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*WeightedUpstreamList)
	if !ok {
		that2, ok := that.(WeightedUpstreamList)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if m.GetWeight() != target.GetWeight() {
		return false
	}

	if h, ok := interface{}(m.GetUpstreamRef()).(equality.Equalizer); ok {
		if !h.Equal(target.GetUpstreamRef()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetUpstreamRef(), target.GetUpstreamRef()) {
			return false
		}
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(equality.Equalizer); ok {
		if !h.Equal(target.GetHeaderManipulation()) {
			return false
		}
	} else {
		if !proto.Equal(m.GetHeaderManipulation(), target.GetHeaderManipulation()) {
			return false
		}
	}

	return true
}
