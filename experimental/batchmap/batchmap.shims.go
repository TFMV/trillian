// Licensed to the Apache Software Foundation (ASF) under one or more
// contributor license agreements.  See the NOTICE file distributed with
// this work for additional information regarding copyright ownership.
// The ASF licenses this file to You under the Apache License, Version 2.0
// (the "License"); you may not use this file except in compliance with
// the License.  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by starcgen. DO NOT EDIT.
// File: batchmap.shims.go

package batchmap

import (
	"fmt"
	"io"
	"reflect"

	// Library imports
	"context"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/exec"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/runtime/graphx/schema"
	"github.com/apache/beam/sdks/v2/go/pkg/beam/core/util/reflectx"
)

func init() {
	runtime.RegisterFunction(entryToNodeHashFn)
	runtime.RegisterFunction(partitionByPrefixLenFn)
	runtime.RegisterFunction(tileToNodeHashFn)
	runtime.RegisterType(reflect.TypeOf((*Entry)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*Entry)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*Tile)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*Tile)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*context.Context)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*context.Context)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*leafShardFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*leafShardFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*nodeHash)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*nodeHash)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*tileHashFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*tileHashFn)(nil)).Elem())
	runtime.RegisterType(reflect.TypeOf((*tileUpdateFn)(nil)).Elem())
	schema.RegisterType(reflect.TypeOf((*tileUpdateFn)(nil)).Elem())
	reflectx.RegisterStructWrapper(reflect.TypeOf((*leafShardFn)(nil)).Elem(), wrapMakerLeafShardFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*tileHashFn)(nil)).Elem(), wrapMakerTileHashFn)
	reflectx.RegisterStructWrapper(reflect.TypeOf((*tileUpdateFn)(nil)).Elem(), wrapMakerTileUpdateFn)
	reflectx.RegisterFunc(reflect.TypeOf((*func(context.Context,[]byte,func(*nodeHash) bool) (*Tile,error))(nil)).Elem(), funcMakerContext۰ContextSliceOfByteIterNodeHashГᏘTileError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(context.Context,[]byte,func(**Tile) bool,func(*nodeHash) bool) (*Tile,error))(nil)).Elem(), funcMakerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError)
	reflectx.RegisterFunc(reflect.TypeOf((*func(nodeHash) ([]byte,nodeHash))(nil)).Elem(), funcMakerNodeHashГSliceOfByteNodeHash)
	reflectx.RegisterFunc(reflect.TypeOf((*func() ())(nil)).Elem(), funcMakerГ)
	reflectx.RegisterFunc(reflect.TypeOf((*func(*Entry) (nodeHash))(nil)).Elem(), funcMakerᏘEntryГNodeHash)
	reflectx.RegisterFunc(reflect.TypeOf((*func(*Tile) (int))(nil)).Elem(), funcMakerᏘTileГInt)
	reflectx.RegisterFunc(reflect.TypeOf((*func(*Tile) (nodeHash))(nil)).Elem(), funcMakerᏘTileГNodeHash)
	exec.RegisterInput(reflect.TypeOf((*func(*nodeHash) (bool))(nil)).Elem(), iterMakerNodeHash)
	exec.RegisterInput(reflect.TypeOf((*func(**Tile) (bool))(nil)).Elem(), iterMakerᏘTile)
}

func wrapMakerLeafShardFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*leafShardFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 nodeHash) ([]byte, nodeHash) { return dfn.ProcessElement(a0) }),
	}
}

func wrapMakerTileHashFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*tileHashFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 context.Context, a1 []byte, a2 func(*nodeHash) bool) (*Tile, error) { return dfn.ProcessElement(a0, a1, a2) }),
		"Setup": reflectx.MakeFunc(func() { dfn.Setup() }),
	}
}

func wrapMakerTileUpdateFn(fn any) map[string]reflectx.Func {
	dfn := fn.(*tileUpdateFn)
	return map[string]reflectx.Func{
		"ProcessElement": reflectx.MakeFunc(func(a0 context.Context, a1 []byte, a2 func(**Tile) bool, a3 func(*nodeHash) bool) (*Tile, error) { return dfn.ProcessElement(a0, a1, a2, a3) }),
		"Setup": reflectx.MakeFunc(func() { dfn.Setup() }),
	}
}

type callerContext۰ContextSliceOfByteIterNodeHashГᏘTileError struct {
	fn func(context.Context,[]byte,func(*nodeHash) bool) (*Tile,error)
}

func funcMakerContext۰ContextSliceOfByteIterNodeHashГᏘTileError(fn any) reflectx.Func {
	f := fn.(func(context.Context,[]byte,func(*nodeHash) bool) (*Tile,error))
	return &callerContext۰ContextSliceOfByteIterNodeHashГᏘTileError{fn: f}
}

func (c *callerContext۰ContextSliceOfByteIterNodeHashГᏘTileError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerContext۰ContextSliceOfByteIterNodeHashГᏘTileError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerContext۰ContextSliceOfByteIterNodeHashГᏘTileError) Call(args []any) []any {
	out0, out1 := c.fn(args[0].(context.Context), args[1].([]byte), args[2].(func(*nodeHash) bool))
	return []any{out0, out1}
}

func (c *callerContext۰ContextSliceOfByteIterNodeHashГᏘTileError) Call3x2(arg0, arg1, arg2 any) (any, any) {
	return c.fn(arg0.(context.Context), arg1.([]byte), arg2.(func(*nodeHash) bool))
}

type callerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError struct {
	fn func(context.Context,[]byte,func(**Tile) bool,func(*nodeHash) bool) (*Tile,error)
}

func funcMakerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError(fn any) reflectx.Func {
	f := fn.(func(context.Context,[]byte,func(**Tile) bool,func(*nodeHash) bool) (*Tile,error))
	return &callerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError{fn: f}
}

func (c *callerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError) Call(args []any) []any {
	out0, out1 := c.fn(args[0].(context.Context), args[1].([]byte), args[2].(func(**Tile) bool), args[3].(func(*nodeHash) bool))
	return []any{out0, out1}
}

func (c *callerContext۰ContextSliceOfByteIterᏘTileIterNodeHashГᏘTileError) Call4x2(arg0, arg1, arg2, arg3 any) (any, any) {
	return c.fn(arg0.(context.Context), arg1.([]byte), arg2.(func(**Tile) bool), arg3.(func(*nodeHash) bool))
}

type callerNodeHashГSliceOfByteNodeHash struct {
	fn func(nodeHash) ([]byte,nodeHash)
}

func funcMakerNodeHashГSliceOfByteNodeHash(fn any) reflectx.Func {
	f := fn.(func(nodeHash) ([]byte,nodeHash))
	return &callerNodeHashГSliceOfByteNodeHash{fn: f}
}

func (c *callerNodeHashГSliceOfByteNodeHash) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerNodeHashГSliceOfByteNodeHash) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerNodeHashГSliceOfByteNodeHash) Call(args []any) []any {
	out0, out1 := c.fn(args[0].(nodeHash))
	return []any{out0, out1}
}

func (c *callerNodeHashГSliceOfByteNodeHash) Call1x2(arg0 any) (any, any) {
	return c.fn(arg0.(nodeHash))
}

type callerГ struct {
	fn func() ()
}

func funcMakerГ(fn any) reflectx.Func {
	f := fn.(func() ())
	return &callerГ{fn: f}
}

func (c *callerГ) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerГ) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerГ) Call(args []any) []any {
	c.fn()
	return []any{}
}

func (c *callerГ) Call0x0() () {
	c.fn()
}

type callerᏘEntryГNodeHash struct {
	fn func(*Entry) (nodeHash)
}

func funcMakerᏘEntryГNodeHash(fn any) reflectx.Func {
	f := fn.(func(*Entry) (nodeHash))
	return &callerᏘEntryГNodeHash{fn: f}
}

func (c *callerᏘEntryГNodeHash) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerᏘEntryГNodeHash) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerᏘEntryГNodeHash) Call(args []any) []any {
	out0 := c.fn(args[0].(*Entry))
	return []any{out0}
}

func (c *callerᏘEntryГNodeHash) Call1x1(arg0 any) (any) {
	return c.fn(arg0.(*Entry))
}

type callerᏘTileГInt struct {
	fn func(*Tile) (int)
}

func funcMakerᏘTileГInt(fn any) reflectx.Func {
	f := fn.(func(*Tile) (int))
	return &callerᏘTileГInt{fn: f}
}

func (c *callerᏘTileГInt) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerᏘTileГInt) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerᏘTileГInt) Call(args []any) []any {
	out0 := c.fn(args[0].(*Tile))
	return []any{out0}
}

func (c *callerᏘTileГInt) Call1x1(arg0 any) (any) {
	return c.fn(arg0.(*Tile))
}

type callerᏘTileГNodeHash struct {
	fn func(*Tile) (nodeHash)
}

func funcMakerᏘTileГNodeHash(fn any) reflectx.Func {
	f := fn.(func(*Tile) (nodeHash))
	return &callerᏘTileГNodeHash{fn: f}
}

func (c *callerᏘTileГNodeHash) Name() string {
	return reflectx.FunctionName(c.fn)
}

func (c *callerᏘTileГNodeHash) Type() reflect.Type {
	return reflect.TypeOf(c.fn)
}

func (c *callerᏘTileГNodeHash) Call(args []any) []any {
	out0 := c.fn(args[0].(*Tile))
	return []any{out0}
}

func (c *callerᏘTileГNodeHash) Call1x1(arg0 any) (any) {
	return c.fn(arg0.(*Tile))
}

type iterNative struct {
	s     exec.ReStream
	fn    any

	// cur is the "current" stream, if any.
	cur exec.Stream
}

func (v *iterNative) Init() error {
	cur, err := v.s.Open()
	if err != nil {
		return err
	}
	v.cur = cur
	return nil
}

func (v *iterNative) Value() any {
	return v.fn
}

func (v *iterNative) Reset() error {
	if err := v.cur.Close(); err != nil {
		return err
	}
	v.cur = nil
	return nil
}

func iterMakerNodeHash(s exec.ReStream) exec.ReusableInput {
	ret := &iterNative{s: s}
	ret.fn = ret.readNodeHash
	return ret
}

func (v *iterNative) readNodeHash(value *nodeHash) bool {
	elm, err := v.cur.Read()
	if err != nil {
		if err == io.EOF {
			return false
		}
		panic(fmt.Sprintf("broken stream: %v", err))
	}
	*value = elm.Elm.(nodeHash)
	return true
}

func iterMakerᏘTile(s exec.ReStream) exec.ReusableInput {
	ret := &iterNative{s: s}
	ret.fn = ret.readᏘTile
	return ret
}

func (v *iterNative) readᏘTile(value **Tile) bool {
	elm, err := v.cur.Read()
	if err != nil {
		if err == io.EOF {
			return false
		}
		panic(fmt.Sprintf("broken stream: %v", err))
	}
	*value = elm.Elm.(*Tile)
	return true
}


// DO NOT MODIFY: GENERATED CODE
