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

import "context"

type Group struct {
	errc   chan error
	cancel context.CancelCauseFunc
	count  int
}

func NewGroup(cancel context.CancelCauseFunc) *Group {
	return &Group{errc: make(chan error, 1), cancel: cancel}
}

func (g *Group) Go(f func() error) {
	g.count++
	go func() {
		g.errc <- f()
	}()
}

func (g *Group) Wait() error {
	var err error
	for range g.count {
		if e := <-g.errc; e != nil && err == nil {
			err = e
			g.cancel(e)
		}
	}

	return err
}
