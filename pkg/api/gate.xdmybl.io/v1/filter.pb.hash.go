// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/xdmybl/gate-type/proto/gate/v1/filter.proto

package v1

import (
	"encoding/binary"
	"errors"
	"fmt"
	"hash"
	"hash/fnv"

	safe_hasher "github.com/solo-io/protoc-gen-ext/pkg/hasher"
	"github.com/solo-io/protoc-gen-ext/pkg/hasher/hashstructure"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = new(hash.Hash64)
	_ = fnv.New64
	_ = hashstructure.Hash
	_ = new(safe_hasher.SafeHasher)
)

// Hash function
func (m *FilterSpec) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.FilterSpec")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetCommonInfo()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("CommonInfo")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetCommonInfo(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("CommonInfo")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.FilterType.(type) {

	case *FilterSpec_Hcm:

		if h, ok := interface{}(m.GetHcm()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("Hcm")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetHcm(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("Hcm")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *FilterSpec_TcpProxy:

		if h, ok := interface{}(m.GetTcpProxy()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("TcpProxy")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetTcpProxy(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("TcpProxy")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpConnectionManager) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.HttpConnectionManager")); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetRouteConfig()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("RouteConfig")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetRouteConfig(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("RouteConfig")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	for _, v := range m.GetHttpFilter() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetMergeSlashes())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSkipXffAppend())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteConfig) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.RouteConfig")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	for _, v := range m.GetVhostLs() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HeaderManipulation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaderManipulation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HeaderManipulation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *VHost) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.VHost")); err != nil {
		return 0, err
	}

	for _, v := range m.GetRouteLs() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	for _, v := range m.GetDomains() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *Route) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.Route")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetId())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetName())); err != nil {
		return 0, err
	}

	if h, ok := interface{}(m.GetMatch()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("Match")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetMatch(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("Match")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	if h, ok := interface{}(m.GetHeaderManipulation()).(safe_hasher.SafeHasher); ok {
		if _, err = hasher.Write([]byte("HeaderManipulation")); err != nil {
			return 0, err
		}
		if _, err = h.Hash(hasher); err != nil {
			return 0, err
		}
	} else {
		if fieldValue, err := hashstructure.Hash(m.GetHeaderManipulation(), nil); err != nil {
			return 0, err
		} else {
			if _, err = hasher.Write([]byte("HeaderManipulation")); err != nil {
				return 0, err
			}
			if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
				return 0, err
			}
		}
	}

	switch m.RouteActionType.(type) {

	case *Route_ForwardAction:

		if h, ok := interface{}(m.GetForwardAction()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("ForwardAction")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetForwardAction(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("ForwardAction")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Route_RedirectAction:

		if h, ok := interface{}(m.GetRedirectAction()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("RedirectAction")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetRedirectAction(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("RedirectAction")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	case *Route_DirectionAction:

		if h, ok := interface{}(m.GetDirectionAction()).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("DirectionAction")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(m.GetDirectionAction(), nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("DirectionAction")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *ForwardAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.ForwardAction")); err != nil {
		return 0, err
	}

	for _, v := range m.GetWeightClusterLs() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTimeout())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RedirectAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.RedirectAction")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetHostRedirect())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetResponseCode())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetSchemeRewriteSpecifier())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStripQuery())
	if err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetPortRedirect())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *DirectionAction) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.DirectionAction")); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetStatusCode())
	if err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetBody())); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *RouteMatch) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.RouteMatch")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPrefix())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPath())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRegex())); err != nil {
		return 0, err
	}

	for _, v := range m.GetHeaders() {

		if h, ok := interface{}(v).(safe_hasher.SafeHasher); ok {
			if _, err = hasher.Write([]byte("")); err != nil {
				return 0, err
			}
			if _, err = h.Hash(hasher); err != nil {
				return 0, err
			}
		} else {
			if fieldValue, err := hashstructure.Hash(v, nil); err != nil {
				return 0, err
			} else {
				if _, err = hasher.Write([]byte("")); err != nil {
					return 0, err
				}
				if err := binary.Write(hasher, binary.LittleEndian, fieldValue); err != nil {
					return 0, err
				}
			}
		}

	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HeaderMatch) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.HeaderMatch")); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetExact())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetPrefix())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetRegex())); err != nil {
		return 0, err
	}

	if _, err = hasher.Write([]byte(m.GetContains())); err != nil {
		return 0, err
	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetInvert())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *WeightCluster) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.WeightCluster")); err != nil {
		return 0, err
	}

	for _, v := range m.GetClusterName() {

		if _, err = hasher.Write([]byte(v)); err != nil {
			return 0, err
		}

	}

	err = binary.Write(hasher, binary.LittleEndian, m.GetTotalWeight())
	if err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *HttpFilter) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.HttpFilter")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}

// Hash function
func (m *TcpProxy) Hash(hasher hash.Hash64) (uint64, error) {
	if m == nil {
		return 0, nil
	}
	if hasher == nil {
		hasher = fnv.New64()
	}
	var err error
	if _, err = hasher.Write([]byte("gate.xdmybl.io.github.com/xdmybl/gate-type/pkg/api/gate.xdmybl.io/v1.TcpProxy")); err != nil {
		return 0, err
	}

	return hasher.Sum64(), nil
}
