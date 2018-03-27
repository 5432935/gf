// Copyright 2018 gf Author(https://gitee.com/johng/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://gitee.com/johng/gf.

package gtype

import (
    "sync"
)

type Bytes struct {
    mu  sync.RWMutex
    val []byte
}

func NewBytes(value...[]byte) *Bytes {
    if len(value) > 0 {
        return &Bytes{val:value[0]}
    }
    return &Bytes{}
}

func (t *Bytes)Set(value []byte) {
    t.mu.Lock()
    t.val = value
    t.mu.Unlock()
}

func (t *Bytes)Val() []byte {
    t.mu.RLock()
    b := t.val
    t.mu.RUnlock()
    return b
}
