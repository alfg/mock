// Automatically generated by MockGen. DO NOT EDIT!
// Source: sample/user.go

package mock_user

import (
	bufio "bufio"
	btz "bytes"
	gomock "github.com/dsymonds/gomock/gomock"
	imp1 "github.com/dsymonds/gomock/sample/imp1"
	renamed2 "github.com/dsymonds/gomock/sample/imp2"
	. "github.com/dsymonds/gomock/sample/imp3"
	imp_four "github.com/dsymonds/gomock/sample/imp4"
	hash "hash"
	io "io"
	http "net/http"
)

// Mock of Index interface
type MockIndex struct {
	ctrl     *gomock.Controller
	recorder *_MockIndexRecorder
}

// Recorder for MockIndex (not exported)
type _MockIndexRecorder struct {
	mock *MockIndex
}

func NewMockIndex(ctrl *gomock.Controller) *MockIndex {
	mock := &MockIndex{ctrl: ctrl}
	mock.recorder = &_MockIndexRecorder{mock}
	return mock
}

func (_m *MockIndex) EXPECT() *_MockIndexRecorder {
	return _m.recorder
}

func (_m *MockIndex) Get(key string) interface{} {
	ret := _m.ctrl.Call(_m, "Get", key)
	ret0, _ := ret[0].(interface{})
	return ret0
}

func (_mr *_MockIndexRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockIndex) GetTwo(key1 string, key2 string) (interface{}, interface{}) {
	ret := _m.ctrl.Call(_m, "GetTwo", key1, key2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(interface{})
	return ret0, ret1
}

func (_mr *_MockIndexRecorder) GetTwo(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetTwo", arg0, arg1)
}

func (_m *MockIndex) Put(key string, value interface{}) {
	_m.ctrl.Call(_m, "Put", key, value)
}

func (_mr *_MockIndexRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0, arg1)
}

func (_m *MockIndex) Summary(buf *btz.Buffer, w io.Writer) {
	_m.ctrl.Call(_m, "Summary", buf, w)
}

func (_mr *_MockIndexRecorder) Summary(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Summary", arg0, arg1)
}

func (_m *MockIndex) Other() hash.Hash {
	ret := _m.ctrl.Call(_m, "Other")
	ret0, _ := ret[0].(hash.Hash)
	return ret0
}

func (_mr *_MockIndexRecorder) Other() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Other")
}

func (_m *MockIndex) Anon(_param0 string) {
	_m.ctrl.Call(_m, "Anon", _param0)
}

func (_mr *_MockIndexRecorder) Anon(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Anon", arg0)
}

func (_m *MockIndex) ForeignOne(_param0 imp1.Imp1) {
	_m.ctrl.Call(_m, "ForeignOne", _param0)
}

func (_mr *_MockIndexRecorder) ForeignOne(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignOne", arg0)
}

func (_m *MockIndex) ForeignTwo(_param0 renamed2.Imp2) {
	_m.ctrl.Call(_m, "ForeignTwo", _param0)
}

func (_mr *_MockIndexRecorder) ForeignTwo(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignTwo", arg0)
}

func (_m *MockIndex) ForeignThree(_param0 Imp3) {
	_m.ctrl.Call(_m, "ForeignThree", _param0)
}

func (_mr *_MockIndexRecorder) ForeignThree(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignThree", arg0)
}

func (_m *MockIndex) ForeignFour(_param0 imp_four.Imp4) {
	_m.ctrl.Call(_m, "ForeignFour", _param0)
}

func (_mr *_MockIndexRecorder) ForeignFour(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignFour", arg0)
}

func (_m *MockIndex) NillableRet() error {
	ret := _m.ctrl.Call(_m, "NillableRet")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockIndexRecorder) NillableRet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NillableRet")
}

func (_m *MockIndex) ConcreteRet() chan<- bool {
	ret := _m.ctrl.Call(_m, "ConcreteRet")
	ret0, _ := ret[0].(chan<- bool)
	return ret0
}

func (_mr *_MockIndexRecorder) ConcreteRet() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConcreteRet")
}

func (_m *MockIndex) Ellip(fmt string, args ...interface{}) {
	_s := []interface{}{fmt}
	for _, _x := range args {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "Ellip", _s...)
}

func (_mr *_MockIndexRecorder) Ellip(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	_s := append([]interface{}{arg0}, arg1...)
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ellip", _s...)
}

func (_m *MockIndex) EllipOnly(_param0 ...string) {
	_s := []interface{}{}
	for _, _x := range _param0 {
		_s = append(_s, _x)
	}
	_m.ctrl.Call(_m, "EllipOnly", _s...)
}

func (_mr *_MockIndexRecorder) EllipOnly(arg0 ...interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EllipOnly", arg0...)
}

func (_m *MockIndex) Ptr(arg *int) {
	_m.ctrl.Call(_m, "Ptr", arg)
}

func (_mr *_MockIndexRecorder) Ptr(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Ptr", arg0)
}

func (_m *MockIndex) Slice(a []int) [3]int {
	ret := _m.ctrl.Call(_m, "Slice", a)
	ret0, _ := ret[0].([3]int)
	return ret0
}

func (_mr *_MockIndexRecorder) Slice(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Slice", arg0)
}

func (_m *MockIndex) Chan(a chan int, b chan<- hash.Hash) {
	_m.ctrl.Call(_m, "Chan", a, b)
}

func (_mr *_MockIndexRecorder) Chan(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Chan", arg0, arg1)
}

func (_m *MockIndex) Func(f func(http.Request) (int, bool)) {
	_m.ctrl.Call(_m, "Func", f)
}

func (_mr *_MockIndexRecorder) Func(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Func", arg0)
}

func (_m *MockIndex) Map(a map[int]hash.Hash) {
	_m.ctrl.Call(_m, "Map", a)
}

func (_mr *_MockIndexRecorder) Map(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Map", arg0)
}

// Mock of Embed interface
type MockEmbed struct {
	ctrl     *gomock.Controller
	recorder *_MockEmbedRecorder
}

// Recorder for MockEmbed (not exported)
type _MockEmbedRecorder struct {
	mock *MockEmbed
}

func NewMockEmbed(ctrl *gomock.Controller) *MockEmbed {
	mock := &MockEmbed{ctrl: ctrl}
	mock.recorder = &_MockEmbedRecorder{mock}
	return mock
}

func (_m *MockEmbed) EXPECT() *_MockEmbedRecorder {
	return _m.recorder
}

func (_m *MockEmbed) RegularMethod() {
	_m.ctrl.Call(_m, "RegularMethod")
}

func (_mr *_MockEmbedRecorder) RegularMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegularMethod")
}

// Methods for embedded interface Embedded

func (_m *MockEmbed) EmbeddedMethod() {
	_m.ctrl.Call(_m, "EmbeddedMethod")
}

func (_mr *_MockEmbedRecorder) EmbeddedMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EmbeddedMethod")
}

// Methods for embedded interface imp1.ForeignEmbedded

func (_m *MockEmbed) ForeignEmbeddedMethod() *bufio.Reader {
	ret := _m.ctrl.Call(_m, "ForeignEmbeddedMethod")
	ret0, _ := ret[0].(*bufio.Reader)
	return ret0
}

func (_mr *_MockEmbedRecorder) ForeignEmbeddedMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ForeignEmbeddedMethod")
}

func (_m *MockEmbed) ImplicitPackage(s string, t imp1.ImpT, st []imp1.ImpT, pt *imp1.ImpT, ct chan imp1.ImpT) {
	_m.ctrl.Call(_m, "ImplicitPackage", s, t, st, pt, ct)
}

func (_mr *_MockEmbedRecorder) ImplicitPackage(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ImplicitPackage", arg0, arg1, arg2, arg3, arg4)
}

// Mock of Embedded interface
type MockEmbedded struct {
	ctrl     *gomock.Controller
	recorder *_MockEmbeddedRecorder
}

// Recorder for MockEmbedded (not exported)
type _MockEmbeddedRecorder struct {
	mock *MockEmbedded
}

func NewMockEmbedded(ctrl *gomock.Controller) *MockEmbedded {
	mock := &MockEmbedded{ctrl: ctrl}
	mock.recorder = &_MockEmbeddedRecorder{mock}
	return mock
}

func (_m *MockEmbedded) EXPECT() *_MockEmbeddedRecorder {
	return _m.recorder
}

func (_m *MockEmbedded) EmbeddedMethod() {
	_m.ctrl.Call(_m, "EmbeddedMethod")
}

func (_mr *_MockEmbeddedRecorder) EmbeddedMethod() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "EmbeddedMethod")
}