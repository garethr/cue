// Copyright 2018 The CUE Authors
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

package cue

// validate returns whether there is any error, recursively.
func validate(ctx *context, v value) *bottom {
	eval := v.evalPartial(ctx)
	if err, ok := eval.(*bottom); ok && err.code != codeIncomplete {
		return eval.(*bottom)
	}
	switch x := eval.(type) {
	case *structLit:
		x = x.expandFields(ctx)
		for i := range x.arcs {
			if err := validate(ctx, x.at(ctx, i)); err != nil {
				return err
			}
		}
	case *list:
		for _, v := range x.a {
			if err := validate(ctx, v); err != nil {
				return err
			}
		}
	}
	return nil
}
