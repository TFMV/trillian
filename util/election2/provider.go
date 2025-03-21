// Copyright 2025 Google LLC. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package election2

import (
	"fmt"
	"sync"
)

var (
	qpMu     sync.RWMutex
	qpByName map[string]NewFactoryFunc
)

// NewFactoryFunc is the signature of a function which can be registered
// to provide instances of an election factory.
type NewFactoryFunc func() (Factory, error)

// RegisterProvider registers a function that provides Manager instances.
func RegisterProvider(name string, qp NewFactoryFunc) error {
	qpMu.Lock()
	defer qpMu.Unlock()

	if qpByName == nil {
		qpByName = make(map[string]NewFactoryFunc)
	}

	_, exists := qpByName[name]
	if exists {
		return fmt.Errorf("election provider %v already registered", name)
	}
	qpByName[name] = qp
	return nil
}

// Providers returns a slice of registered quota provider names.
func Providers() []string {
	qpMu.RLock()
	defer qpMu.RUnlock()

	r := []string{}
	for k := range qpByName {
		r = append(r, k)
	}

	return r
}

// NewProvider returns a Manager implementation.
func NewProvider(name string) (Factory, error) {
	qpMu.RLock()
	defer qpMu.RUnlock()

	f, exists := qpByName[name]
	if !exists {
		return nil, fmt.Errorf("unknown election provider: %v", name)
	}
	return f()
}
