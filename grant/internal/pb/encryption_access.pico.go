// Code generated by protoc-gen-pico. DO NOT EDIT.
// source: encryption_access.proto
//
// versions:
//     protoc-gen-pico: v0.0.3
//     protoc:          v4.25.1

package pb

import (
	picobuf "storj.io/picobuf"
)

type EncryptionAccess struct {
	DefaultKey                  []byte                         `json:"default_key,omitempty"`
	StoreEntries                []*EncryptionAccess_StoreEntry `json:"store_entries,omitempty"`
	DefaultPathCipher           CipherSuite                    `json:"default_path_cipher,omitempty"`
	DefaultEncryptionParameters *EncryptionParameters          `json:"default_encryption_parameters,omitempty"`
}

func (m *EncryptionAccess) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Bytes(1, &m.DefaultKey)
	for _, x := range m.StoreEntries {
		c.AlwaysMessage(2, x.Encode)
	}
	c.Int32(3, (*int32)(&m.DefaultPathCipher))
	c.Message(4, m.DefaultEncryptionParameters.Encode)
	return true
}

func (m *EncryptionAccess) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Bytes(1, &m.DefaultKey)
	c.RepeatedMessage(2, func(c *picobuf.Decoder) {
		x := new(EncryptionAccess_StoreEntry)
		c.Loop(x.Decode)
		m.StoreEntries = append(m.StoreEntries, x)
	})
	c.Int32(3, (*int32)(&m.DefaultPathCipher))
	c.Message(4, func(c *picobuf.Decoder) {
		if m.DefaultEncryptionParameters == nil {
			m.DefaultEncryptionParameters = new(EncryptionParameters)
		}
		m.DefaultEncryptionParameters.Decode(c)
	})
}

type EncryptionAccess_StoreEntry struct {
	Bucket               []byte                `json:"bucket,omitempty"`
	UnencryptedPath      []byte                `json:"unencrypted_path,omitempty"`
	EncryptedPath        []byte                `json:"encrypted_path,omitempty"`
	Key                  []byte                `json:"key,omitempty"`
	PathCipher           CipherSuite           `json:"path_cipher,omitempty"`
	EncryptionParameters *EncryptionParameters `json:"encryption_parameters,omitempty"`
}

func (m *EncryptionAccess_StoreEntry) Encode(c *picobuf.Encoder) bool {
	if m == nil {
		return false
	}
	c.Bytes(1, &m.Bucket)
	c.Bytes(2, &m.UnencryptedPath)
	c.Bytes(3, &m.EncryptedPath)
	c.Bytes(4, &m.Key)
	c.Int32(5, (*int32)(&m.PathCipher))
	c.Message(6, m.EncryptionParameters.Encode)
	return true
}

func (m *EncryptionAccess_StoreEntry) Decode(c *picobuf.Decoder) {
	if m == nil {
		return
	}
	c.Bytes(1, &m.Bucket)
	c.Bytes(2, &m.UnencryptedPath)
	c.Bytes(3, &m.EncryptedPath)
	c.Bytes(4, &m.Key)
	c.Int32(5, (*int32)(&m.PathCipher))
	c.Message(6, func(c *picobuf.Decoder) {
		if m.EncryptionParameters == nil {
			m.EncryptionParameters = new(EncryptionParameters)
		}
		m.EncryptionParameters.Decode(c)
	})
}
