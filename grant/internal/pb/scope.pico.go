// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: scope.proto
//
// versions:
//     protoc-gen-pico: v0.0.4
//     protoc:          v6.30.2

package pb

import (
	picobuf "storj.io/picobuf"
)

type Scope struct {
	SatelliteAddr    string            `json:"satellite_addr,omitempty"`
	ApiKey           []byte            `json:"api_key,omitempty"`
	EncryptionAccess *EncryptionAccess `json:"encryption_access,omitempty"`
}

func (m *Scope) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.String(1, &m.SatelliteAddr)
	c.Bytes(2, &m.ApiKey)
	c.Message(3, m.EncryptionAccess.Encode)
	return true
}

func (m *Scope) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.String(1, &m.SatelliteAddr)
	c.Bytes(2, &m.ApiKey)
	c.Message(3, func(c *picobuf.Decoder) {
		if m.EncryptionAccess == nil {
			m.EncryptionAccess = new(EncryptionAccess)
		}
		m.EncryptionAccess.Decode(c)
	})
}
