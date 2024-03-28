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

	"fillmore-labs.com/blog/structured/pkg/task"
)

func doWork(ctx context.Context) error {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	var g int
	errc := make(chan error)

	g++
	go func() {
		errc <- task.Task(ctx, "task1", processingTime/3, nil)
	}()

	g++
	go func() {
		errc <- task.Task(ctx, "task2", processingTime/2, errFail)
	}()

	g++
	go func() {
		errc <- task.Task(ctx, "task3", processingTime, nil)
	}()

	var err error
	for range g {
		if e := <-errc; e != nil && err == nil {
			err = e
			cancel(err)
		}
	}

	return err
}
