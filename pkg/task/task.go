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

package task

import (
	"context"
	"fmt"
	"time"
)

func Task(ctx context.Context, name string, processingTime time.Duration, result error) error {
	ready := time.NewTimer(processingTime)

	select {
	case <-ctx.Done():
		ready.Stop()
		fmt.Println(name, ctx.Err())

		return fmt.Errorf("%s canceled: %w", name, ctx.Err())

	case <-ready.C:
		fmt.Println(name, result)
	}

	return result
}
