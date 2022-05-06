// Copyright 2022 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package unary

import (
	"fmt"

	"github.com/matrixorigin/matrixone/pkg/builtin"
	"github.com/matrixorigin/matrixone/pkg/container/nulls"
	"github.com/matrixorigin/matrixone/pkg/container/types"
	"github.com/matrixorigin/matrixone/pkg/container/vector"
	"github.com/matrixorigin/matrixone/pkg/encoding"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/extend"
	"github.com/matrixorigin/matrixone/pkg/sql/colexec/extend/overload"
	"github.com/matrixorigin/matrixone/pkg/vectorize/notempty"
	"github.com/matrixorigin/matrixone/pkg/vm/process"
)

func init() {
	extend.FunctionRegistry["notempty"] = builtin.NotEmpty

	// register function args and returns
	overload.AppendFunctionRets(builtin.NotEmpty, []types.T{types.T_char}, types.T_uint8)
	overload.AppendFunctionRets(builtin.NotEmpty, []types.T{types.T_varchar}, types.T_uint8)

	extend.UnaryReturnTypes[builtin.NotEmpty] = func(extend extend.Extend) types.T {
		return getUnaryReturnType(builtin.NotEmpty, extend) // register a get return type function
	}

	extend.UnaryStrings[builtin.NotEmpty] = func(e extend.Extend) string {
		return fmt.Sprintf("notempty(%s)", e) // define a stringify function
	}

	overload.OpTypes[builtin.NotEmpty] = overload.Unary // register function type

	overload.UnaryOps[builtin.NotEmpty] = []*overload.UnaryOp{
		{
			Typ:        types.T_varchar,
			ReturnType: types.T_uint8,
			Fn:         notEmptyFn,
		},
		{
			Typ:        types.T_char,
			ReturnType: types.T_uint8,
			Fn:         notEmptyFn,
		},
	}
}

func notEmptyFn(origVec *vector.Vector, proc *process.Process, _ bool) (*vector.Vector, error) {
	col := origVec.Col.(*types.Bytes)
	resultLen := len(col.Lengths)
	resultVec, err := process.Get(proc, int64(resultLen), types.Type{Oid: types.T_uint8, Size: 1})
	if err != nil {
		return nil, err
	}
	result := encoding.DecodeUint8Slice(resultVec.Data)
	result = result[:resultLen]
	resultVec.Col = result
	// the new vector's nulls are the same as the original vector
	nulls.Set(resultVec.Nsp, origVec.Nsp)
	vector.SetCol(resultVec, notempty.NotEmpty(col, result))
	return resultVec, nil
}
