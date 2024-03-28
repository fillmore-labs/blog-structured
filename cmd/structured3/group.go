// Copyright 2024 Oliver Eikemeier. All Rights Reserved.
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
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"sync"
)

type Group struct {
	err    error
	cancel context.CancelCauseFunc
	once   sync.Once
	wg     sync.WaitGroup
}

func NewGroup(cancel context.CancelCauseFunc) *Group {
	return &Group{cancel: cancel}
}

func (g *Group) Do(fn func() error) {
	err := fn()
	if err == nil {
		return
	}
	g.once.Do(func() {
		g.err = err
		if g.cancel != nil {
			g.cancel(err)
		}
	})
}

func (g *Group) Go(fn func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		g.Do(fn)
	}()
}

func (g *Group) Wait() error {
	g.wg.Wait()

	return g.err
}
