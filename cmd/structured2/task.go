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

	"fillmore-labs.com/blog/structured/pkg/task"
)

func doWork(ctx context.Context) error {
	ctx, cancel := context.WithCancelCause(ctx)
	defer cancel(nil)

	var firstErr error
	var once sync.Once
	setErr := func(err error) {
		if err == nil {
			return
		}
		once.Do(func() {
			firstErr = err
			cancel(err)
		})
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := task.Task(ctx, "task1", processingTime/3, nil)
		setErr(err)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		err := task.Task(ctx, "task2", processingTime/2, errFail)
		setErr(err)
	}()

	err := task.Task(ctx, "task3", processingTime, nil)
	setErr(err)

	wg.Wait()

	return firstErr
}
