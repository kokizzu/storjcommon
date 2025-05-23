// Copyright (C) 2019 Storj Labs, Inc.
// See LICENSE for copying information.

package encryption

import (
	"encoding/base64"
	"strings"

	"github.com/zeebo/errs"

	"storj.io/common/internal/hmacsha512"
	"storj.io/common/paths"
	"storj.io/common/storj"
)

var (
	emptyComponentPrefix    = byte('\x01')
	notEmptyComponentPrefix = byte('\x02')
	emptyComponent          = []byte{emptyComponentPrefix}

	escapeSlash = byte('\x2e')
	escapeFF    = byte('\xfe')
	escape01    = byte('\x01')

	// ErrMissingEncryptionBase is caused by trying to encrypt a path without
	// having the appropriate root in the encryption store.
	ErrMissingEncryptionBase = errs.Class("missing encryption base")

	// ErrMissingDecryptionBase is caused by trying to decrypt a path without
	// having the appropriate root in the encryption store.
	ErrMissingDecryptionBase = errs.Class("missing decryption base")
)

type pathBuilder struct {
	i int
	b strings.Builder
}

func (p *pathBuilder) append(x string) {
	if p.i > 0 {
		_ = p.b.WriteByte('/')
	}
	_, _ = p.b.WriteString(x)
	p.i++
}

func (p *pathBuilder) String() string                 { return p.b.String() }
func (p *pathBuilder) Encrypted() paths.Encrypted     { return paths.NewEncrypted(p.String()) }
func (p *pathBuilder) Unencrypted() paths.Unencrypted { return paths.NewUnencrypted(p.String()) }

// EncryptPathWithStoreCipher encrypts the path looking up keys and the cipher from the
// provided store and bucket.
func EncryptPathWithStoreCipher(bucket string, path paths.Unencrypted, store *Store) (
	encPath paths.Encrypted, err error) {

	return encryptPath(bucket, path, nil, store)
}

// EncryptPrefixWithStoreCipher encrypts the prefix using the provided cipher and looking up keys from the
// provided store and bucket. Because it is a prefix, it does not assume there is an empty component
// at the end of a path like "foo/bar/".
func EncryptPrefixWithStoreCipher(bucket string, path paths.Unencrypted, store *Store) (
	encPath paths.Encrypted, err error) {

	raw := path.Raw()
	hasTrailing := strings.HasSuffix(raw, "/")
	if hasTrailing {
		path = paths.NewUnencrypted(raw[:len(raw)-1])
	}
	encPath, err = encryptPath(bucket, path, nil, store)
	if err != nil {
		return encPath, err
	}
	if hasTrailing {
		encPath = paths.NewEncrypted(encPath.Raw() + "/")
	}
	return encPath, nil
}

// EncryptPath encrypts the path using the provided cipher and looking up keys from the
// provided store and bucket.
func EncryptPath(bucket string, path paths.Unencrypted, pathCipher storj.CipherSuite, store *Store) (
	encPath paths.Encrypted, err error) {

	return encryptPath(bucket, path, &pathCipher, store)
}

func encryptPath(bucket string, path paths.Unencrypted, pathCipher *storj.CipherSuite, store *Store) (
	encPath paths.Encrypted, err error) {

	// Invalid paths map to invalid paths
	if !path.Valid() {
		return paths.Encrypted{}, nil
	}

	_, remaining, base := store.LookupUnencrypted(bucket, path)
	if base == nil {
		return paths.Encrypted{}, ErrMissingEncryptionBase.New("%q/%q", bucket, path)
	}

	if pathCipher == nil {
		pathCipher = &base.PathCipher
	}

	// if we're using the default base (meaning the default key), we need
	// to include the bucket name in the path derivation.
	key := &base.Key
	if base.Default {
		key, err = derivePathKeyComponent(key, bucket)
		if err != nil {
			return paths.Encrypted{}, errs.Wrap(err)
		}
	}

	encrypted, err := EncryptIterator(remaining, *pathCipher, key)
	if err != nil {
		return paths.Encrypted{}, errs.Wrap(err)
	}

	var pb pathBuilder
	if base.Encrypted.Valid() {
		pb.append(base.Encrypted.Raw())
	}
	if !remaining.Done() {
		pb.append(encrypted)
	}

	return pb.Encrypted(), nil
}

// EncryptIterator encrypts the iterator using the provided key directly, returning the
// joined path. EncryptPath should be preferred if possible.
func EncryptIterator(iter paths.Iterator, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	var pb pathBuilder
	for !iter.Done() {
		component := iter.Next()
		encComponent, err := encryptPathComponent(component, cipher, key)
		if err != nil {
			return "", errs.Wrap(err)
		}
		key, err = derivePathKeyComponent(key, component)
		if err != nil {
			return "", errs.Wrap(err)
		}
		pb.append(encComponent)
	}
	return pb.String(), nil
}

// EncryptPathRaw encrypts the path using the provided key directly. EncryptPath should be
// preferred if possible.
func EncryptPathRaw(raw string, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	return EncryptIterator(paths.NewIterator(raw), cipher, key)
}

// DecryptPathWithStoreCipher decrypts the path looking up keys and the cipher from the
// provided store and bucket.
func DecryptPathWithStoreCipher(bucket string, path paths.Encrypted, store *Store) (
	encPath paths.Unencrypted, err error) {

	return decryptPath(bucket, path, nil, store)
}

// DecryptPath decrypts the path using the provided cipher and looking up keys from the
// provided store and bucket.
func DecryptPath(bucket string, path paths.Encrypted, pathCipher storj.CipherSuite, store *Store) (
	encPath paths.Unencrypted, err error) {

	return decryptPath(bucket, path, &pathCipher, store)
}

func decryptPath(bucket string, path paths.Encrypted, pathCipher *storj.CipherSuite, store *Store) (
	encPath paths.Unencrypted, err error) {

	// Invalid paths map to invalid paths
	if !path.Valid() {
		return paths.Unencrypted{}, nil
	}

	_, remaining, base := store.LookupEncrypted(bucket, path)
	if base == nil {
		return paths.Unencrypted{}, ErrMissingDecryptionBase.New("%q/%q", bucket, path)
	}

	if pathCipher == nil {
		pathCipher = &base.PathCipher
	}

	// if we're using the default base (meaning the default key), we need
	// to include the bucket name in the path derivation.
	key := &base.Key
	if base.Default {
		key, err = derivePathKeyComponent(key, bucket)
		if err != nil {
			return paths.Unencrypted{}, errs.Wrap(err)
		}
	}

	decrypted, err := DecryptIterator(remaining, *pathCipher, key)
	if err != nil {
		return paths.Unencrypted{}, errs.Wrap(err)
	}

	var pb pathBuilder
	if base.Unencrypted.Valid() {
		pb.append(base.Unencrypted.Raw())
	}
	if !remaining.Done() {
		pb.append(decrypted)
	}

	return pb.Unencrypted(), nil
}

// DecryptIterator decrypts the iterator using the provided key directly, returning the
// joined path. DecryptPath should be preferred if possible.
func DecryptIterator(iter paths.Iterator, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	var pb pathBuilder
	for !iter.Done() {
		component := iter.Next()
		unencComponent, err := decryptPathComponent(component, cipher, key)
		if err != nil {
			return "", errs.Wrap(err)
		}
		key, err = derivePathKeyComponent(key, unencComponent)
		if err != nil {
			return "", errs.Wrap(err)
		}
		pb.append(unencComponent)
	}
	return pb.String(), nil
}

// DecryptPathRaw decrypts the path using the provided key directly. DecryptPath should be
// preferred if possible.
func DecryptPathRaw(raw string, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	return DecryptIterator(paths.NewIterator(raw), cipher, key)
}

// DeriveContentKey returns the content key for the passed in path by looking up
// the appropriate base key from the store and bucket and deriving the rest.
func DeriveContentKey(bucket string, path paths.Unencrypted, store *Store) (key *storj.Key, err error) {
	key, err = DerivePathKey(bucket, path, store)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	key, err = DeriveKey(key, "content")
	return key, errs.Wrap(err)
}

// DerivePathKey returns the path key for the passed in path by looking up the
// appropriate base key from the store and bucket and deriving the rest.
func DerivePathKey(bucket string, path paths.Unencrypted, store *Store) (key *storj.Key, err error) {
	_, remaining, base := store.LookupUnencrypted(bucket, path)
	if base == nil {
		return nil, ErrMissingEncryptionBase.New("%q/%q", bucket, path)
	}

	// If asking for the key at the bucket, do that and return.
	if !path.Valid() {
		// if we're using the default base (meaning the default key), we need
		// to include the bucket name in the path derivation.
		key = &base.Key
		if base.Default {
			key, err = derivePathKeyComponent(&base.Key, bucket)
			if err != nil {
				return nil, errs.Wrap(err)
			}
		}
		return key, nil
	}

	// if we're using the default base (meaning the default key), we need
	// to include the bucket name in the path derivation.
	key = &base.Key
	if base.Default {
		key, err = derivePathKeyComponent(key, bucket)
		if err != nil {
			return nil, errs.Wrap(err)
		}
	}

	for !remaining.Done() {
		key, err = derivePathKeyComponent(key, remaining.Next())
		if err != nil {
			return nil, errs.Wrap(err)
		}
	}
	return key, nil
}

// derivePathKeyComponent derives a new key from the provided one using the component. It
// should be preferred over DeriveKey when adding path components as it performs the
// necessary transformation to the component.
func derivePathKeyComponent(key *storj.Key, component string) (*storj.Key, error) {
	return DeriveKey(key, "path:"+component)
}

// encryptPathComponent encrypts a single path component with the provided cipher and key.
func encryptPathComponent(comp string, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	if cipher == storj.EncNull {
		return comp, nil
	}

	if cipher == storj.EncNullBase64URL {
		decoded, err := base64.URLEncoding.DecodeString(comp)
		if err != nil {
			return "", Error.New("invalid base64 data: %v", err)
		}
		return string(decoded), nil
	}

	// derive the key for the next path component. this is so that
	// every encrypted component has a unique nonce.
	derivedKey, err := derivePathKeyComponent(key, comp)
	if err != nil {
		return "", err
	}

	// use the derived key to derive the nonce
	mac := hmacsha512.New(derivedKey[:])
	mac.Write([]byte("nonce"))
	nonce := new(storj.Nonce)
	sum := mac.SumAndReset()
	copy(nonce[:], sum[:])

	// encrypt the path components with the parent's key and the derived nonce
	cipherText, err := Encrypt([]byte(comp), cipher, key, nonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	nonceSize := storj.NonceSize
	if cipher == storj.EncAESGCM {
		nonceSize = AESGCMNonceSize
	}

	// keep the nonce together with the cipher text
	return string(encodeSegment(append(nonce[:nonceSize], cipherText...))), nil
}

// decryptPathComponent decrypts a single path component with the provided cipher and key.
func decryptPathComponent(comp string, cipher storj.CipherSuite, key *storj.Key) (string, error) {
	if comp == "" {
		return "", nil
	}

	if cipher == storj.EncNull {
		return comp, nil
	}

	if cipher == storj.EncNullBase64URL {
		return base64.URLEncoding.EncodeToString([]byte(comp)), nil
	}

	data, err := decodeSegment([]byte(comp))
	if err != nil {
		return "", Error.Wrap(err)
	}

	nonceSize := storj.NonceSize
	if cipher == storj.EncAESGCM {
		nonceSize = AESGCMNonceSize
	}
	if len(data) < nonceSize || nonceSize < 0 {
		return "", errs.New("component did not contain enough nonce bytes")
	}

	// extract the nonce from the cipher text
	nonce := new(storj.Nonce)
	copy(nonce[:], data[:nonceSize])

	decrypted, err := Decrypt(data[nonceSize:], cipher, key, nonce)
	if err != nil {
		return "", Error.Wrap(err)
	}

	return string(decrypted), nil
}

// encodeSegment encodes segment according to specific rules
// The empty path component is encoded as `\x01`
// Any other path component is encoded as `\x02 + escape(component)`
//
// `\x2e` escapes to `\x2e\x01`
// `\x2f` escapes to `\x2e\x02`
// `\xfe` escapes to `\xfe\x01`
// `\xff` escapes to `\xfe\x02`
// `\x00` escapes to `\x01\x01`
// `\x01` escapes to `\x01\x02
// for more details see docs/design/path-component-encoding.md.
func encodeSegment(segment []byte) []byte {
	if len(segment) == 0 {
		return emptyComponent
	}

	result := make([]byte, 0, len(segment)*2+1)
	result = append(result, notEmptyComponentPrefix)
	for i := range segment {
		switch segment[i] {
		case escapeSlash:
			result = append(result, []byte{escapeSlash, 1}...)
		case escapeSlash + 1:
			result = append(result, []byte{escapeSlash, 2}...)
		case escapeFF:
			result = append(result, []byte{escapeFF, 1}...)
		case escapeFF + 1:
			result = append(result, []byte{escapeFF, 2}...)
		case escape01 - 1:
			result = append(result, []byte{escape01, 1}...)
		case escape01:
			result = append(result, []byte{escape01, 2}...)
		default:
			result = append(result, segment[i])
		}
	}
	return result
}

func decodeSegment(segment []byte) ([]byte, error) {
	err := validateEncodedSegment(segment)
	if err != nil {
		return []byte{}, err
	}
	if segment[0] == emptyComponentPrefix {
		return []byte{}, nil
	}

	currentIndex := 0
	for i := 1; i < len(segment); i++ {
		switch {
		case i == len(segment)-1:
			segment[currentIndex] = segment[i]
		case segment[i] == escapeSlash || segment[i] == escapeFF:
			segment[currentIndex] = segment[i] + segment[i+1] - 1
			i++
		case segment[i] == escape01:
			segment[currentIndex] = segment[i+1] - 1
			i++
		default:
			segment[currentIndex] = segment[i]
		}
		currentIndex++
	}
	return segment[:currentIndex], nil
}

// validateEncodedSegment checks if:
// * The last byte/sequence is not in {escape1, escape2, escape3}.
// * Any byte after an escape character is \x01 or \x02.
// * It does not contain any characters in {\x00, \xff, \x2f}.
// * It is non-empty.
// * It begins with a character in {\x01, \x02}.
func validateEncodedSegment(segment []byte) error {
	switch {
	case len(segment) == 0:
		return ErrDecryptFailed.New("encoded segment cannot be empty")
	case segment[0] != emptyComponentPrefix && segment[0] != notEmptyComponentPrefix:
		return ErrDecryptFailed.New("invalid segment prefix")
	case segment[0] == emptyComponentPrefix && len(segment) > 1:
		return ErrDecryptFailed.New("segment encoded as empty but contains data")
	case segment[0] == notEmptyComponentPrefix && len(segment) == 1:
		return ErrDecryptFailed.New("segment encoded as not empty but doesn't contain data")
	}

	if len(segment) == 1 {
		return nil
	}

	index := 1
	for ; index < len(segment)-1; index++ {
		if isEscapeByte(segment[index]) {
			if segment[index+1] == 1 || segment[index+1] == 2 {
				index++
				continue
			}
			return ErrDecryptFailed.New("invalid escape sequence")
		}
		if isDisallowedByte(segment[index]) {
			return ErrDecryptFailed.New("invalid character in segment")
		}
	}
	if index == len(segment)-1 {
		if isEscapeByte(segment[index]) {
			return ErrDecryptFailed.New("invalid escape sequence")
		}
		if isDisallowedByte(segment[index]) {
			return ErrDecryptFailed.New("invalid character")
		}
	}

	return nil
}

func isEscapeByte(b byte) bool {
	return b == escapeSlash || b == escapeFF || b == escape01
}

func isDisallowedByte(b byte) bool {
	return b == 0 || b == '\xff' || b == '/'
}
