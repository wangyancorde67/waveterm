// Copyright 2024, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

// Package waveobj defines the core object model for WaveTerm.
// All persistent objects in the system implement the WaveObj interface.
package waveobj

import (
	"fmt"
	"reflect"
	"strings"
)

// ORef is an object reference, combining a type string and a unique ID.
type ORef struct {
	OType string `json:"otype"`
	OID   string `json:"oid"`
}

// String returns a human-readable representation of the ORef.
func (r ORef) String() string {
	return fmt.Sprintf("%s:%s", r.OType, r.OID)
}

// IsEmpty returns true if the ORef has no type or ID set.
func (r ORef) IsEmpty() bool {
	return r.OType == "" && r.OID == ""
}

// ParseORef parses a string of the form "otype:oid" into an ORef.
func ParseORef(s string) (ORef, error) {
	parts := strings.SplitN(s, ":", 2)
	if len(parts) != 2 {
		return ORef{}, fmt.Errorf("invalid oref %q: expected otype:oid", s)
	}
	if parts[0] == "" {
		return ORef{}, fmt.Errorf("invalid oref %q: otype cannot be empty", s)
	}
	if parts[1] == "" {
		return ORef{}, fmt.Errorf("invalid oref %q: oid cannot be empty", s)
	}
	return ORef{OType: parts[0], OID: parts[1]}, nil
}

// WaveObj is the interface that all persistent WaveTerm objects must implement.
type WaveObj interface {
	// GetOType returns the object type string (e.g. "workspace", "block").
	GetOType() string
}

// GetOID returns the OID field of a WaveObj using reflection.
// The struct must have a field tagged with `json:"oid"`.
func GetOID(obj WaveObj) string {
	v := reflect.ValueOf(obj)
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return ""
	}
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get("json")
		tagName := strings.SplitN(tag, ",", 2)[0]
		if tagName == "oid" {
			return v.Field(i).String()
		}
	}
	return ""
}

// ToORef converts a WaveObj to an ORef.
func ToORef(obj WaveObj) ORef {
	return ORef{
		OType: obj.GetOType(),
		OID:   GetOID(obj),
	}
}

// MetaMapType is a generic map for storing arbitrary metadata on wave objects.
type MetaMapType map[string]any

// GetMeta retrieves a typed value from a MetaMapType by key.
// Returns the zero value of T if the key is absent or the type does not match.
func GetMeta[T any](m MetaMapType, key string) T {
	var zero T
	if m == nil {
		return zero
	}
	v, ok := m[key]
	if !ok {
		return zero
	}
	typed, ok := v.(T)
	if !ok {
		return zero
	}
	return typed
}
