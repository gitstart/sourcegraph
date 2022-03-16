// Code generated by go-mockgen 1.1.4; DO NOT EDIT.

package insights

import (
	"context"
	"sync"
)

// MockLoader is a mock implementation of the Loader interface (from the
// package github.com/sourcegraph/sourcegraph/internal/insights) used for
// unit testing.
type MockLoader struct {
	// LoadAllFunc is an instance of a mock function object controlling the
	// behavior of the method LoadAll.
	LoadAllFunc *LoaderLoadAllFunc
	// LoadDashboardsFunc is an instance of a mock function object
	// controlling the behavior of the method LoadDashboards.
	LoadDashboardsFunc *LoaderLoadDashboardsFunc
}

// NewMockLoader creates a new mock of the Loader interface. All methods
// return zero values for all results, unless overwritten.
func NewMockLoader() *MockLoader {
	return &MockLoader{
		LoadAllFunc: &LoaderLoadAllFunc{
			defaultHook: func(context.Context) ([]SearchInsight, error) {
				return nil, nil
			},
		},
		LoadDashboardsFunc: &LoaderLoadDashboardsFunc{
			defaultHook: func(context.Context) ([]SettingDashboard, error) {
				return nil, nil
			},
		},
	}
}

// NewStrictMockLoader creates a new mock of the Loader interface. All
// methods panic on invocation, unless overwritten.
func NewStrictMockLoader() *MockLoader {
	return &MockLoader{
		LoadAllFunc: &LoaderLoadAllFunc{
			defaultHook: func(context.Context) ([]SearchInsight, error) {
				panic("unexpected invocation of MockLoader.LoadAll")
			},
		},
		LoadDashboardsFunc: &LoaderLoadDashboardsFunc{
			defaultHook: func(context.Context) ([]SettingDashboard, error) {
				panic("unexpected invocation of MockLoader.LoadDashboards")
			},
		},
	}
}

// NewMockLoaderFrom creates a new mock of the MockLoader interface. All
// methods delegate to the given implementation, unless overwritten.
func NewMockLoaderFrom(i Loader) *MockLoader {
	return &MockLoader{
		LoadAllFunc: &LoaderLoadAllFunc{
			defaultHook: i.LoadAll,
		},
		LoadDashboardsFunc: &LoaderLoadDashboardsFunc{
			defaultHook: i.LoadDashboards,
		},
	}
}

// LoaderLoadAllFunc describes the behavior when the LoadAll method of the
// parent MockLoader instance is invoked.
type LoaderLoadAllFunc struct {
	defaultHook func(context.Context) ([]SearchInsight, error)
	hooks       []func(context.Context) ([]SearchInsight, error)
	history     []LoaderLoadAllFuncCall
	mutex       sync.Mutex
}

// LoadAll delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockLoader) LoadAll(v0 context.Context) ([]SearchInsight, error) {
	r0, r1 := m.LoadAllFunc.nextHook()(v0)
	m.LoadAllFunc.appendCall(LoaderLoadAllFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the LoadAll method of
// the parent MockLoader instance is invoked and the hook queue is empty.
func (f *LoaderLoadAllFunc) SetDefaultHook(hook func(context.Context) ([]SearchInsight, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// LoadAll method of the parent MockLoader instance invokes the hook at the
// front of the queue and discards it. After the queue is empty, the default
// hook function is invoked for any future action.
func (f *LoaderLoadAllFunc) PushHook(hook func(context.Context) ([]SearchInsight, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LoaderLoadAllFunc) SetDefaultReturn(r0 []SearchInsight, r1 error) {
	f.SetDefaultHook(func(context.Context) ([]SearchInsight, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LoaderLoadAllFunc) PushReturn(r0 []SearchInsight, r1 error) {
	f.PushHook(func(context.Context) ([]SearchInsight, error) {
		return r0, r1
	})
}

func (f *LoaderLoadAllFunc) nextHook() func(context.Context) ([]SearchInsight, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LoaderLoadAllFunc) appendCall(r0 LoaderLoadAllFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LoaderLoadAllFuncCall objects describing
// the invocations of this function.
func (f *LoaderLoadAllFunc) History() []LoaderLoadAllFuncCall {
	f.mutex.Lock()
	history := make([]LoaderLoadAllFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LoaderLoadAllFuncCall is an object that describes an invocation of method
// LoadAll on an instance of MockLoader.
type LoaderLoadAllFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []SearchInsight
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LoaderLoadAllFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LoaderLoadAllFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// LoaderLoadDashboardsFunc describes the behavior when the LoadDashboards
// method of the parent MockLoader instance is invoked.
type LoaderLoadDashboardsFunc struct {
	defaultHook func(context.Context) ([]SettingDashboard, error)
	hooks       []func(context.Context) ([]SettingDashboard, error)
	history     []LoaderLoadDashboardsFuncCall
	mutex       sync.Mutex
}

// LoadDashboards delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockLoader) LoadDashboards(v0 context.Context) ([]SettingDashboard, error) {
	r0, r1 := m.LoadDashboardsFunc.nextHook()(v0)
	m.LoadDashboardsFunc.appendCall(LoaderLoadDashboardsFuncCall{v0, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the LoadDashboards
// method of the parent MockLoader instance is invoked and the hook queue is
// empty.
func (f *LoaderLoadDashboardsFunc) SetDefaultHook(hook func(context.Context) ([]SettingDashboard, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// LoadDashboards method of the parent MockLoader instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *LoaderLoadDashboardsFunc) PushHook(hook func(context.Context) ([]SettingDashboard, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *LoaderLoadDashboardsFunc) SetDefaultReturn(r0 []SettingDashboard, r1 error) {
	f.SetDefaultHook(func(context.Context) ([]SettingDashboard, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *LoaderLoadDashboardsFunc) PushReturn(r0 []SettingDashboard, r1 error) {
	f.PushHook(func(context.Context) ([]SettingDashboard, error) {
		return r0, r1
	})
}

func (f *LoaderLoadDashboardsFunc) nextHook() func(context.Context) ([]SettingDashboard, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *LoaderLoadDashboardsFunc) appendCall(r0 LoaderLoadDashboardsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of LoaderLoadDashboardsFuncCall objects
// describing the invocations of this function.
func (f *LoaderLoadDashboardsFunc) History() []LoaderLoadDashboardsFuncCall {
	f.mutex.Lock()
	history := make([]LoaderLoadDashboardsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// LoaderLoadDashboardsFuncCall is an object that describes an invocation of
// method LoadDashboards on an instance of MockLoader.
type LoaderLoadDashboardsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 []SettingDashboard
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c LoaderLoadDashboardsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c LoaderLoadDashboardsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
