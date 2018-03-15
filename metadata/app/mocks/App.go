// Copyright 2018 The kubecfg authors
//
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.

// Code generated by mockery v1.0.0
package mocks

import afero "github.com/spf13/afero"
import app "github.com/ksonnet/ksonnet/metadata/app"
import mock "github.com/stretchr/testify/mock"

// App is an autogenerated mock type for the App type
type App struct {
	mock.Mock
}

// AddEnvironment provides a mock function with given fields: name, k8sSpecFlag, spec
func (_m *App) AddEnvironment(name string, k8sSpecFlag string, spec *app.EnvironmentSpec) error {
	ret := _m.Called(name, k8sSpecFlag, spec)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, *app.EnvironmentSpec) error); ok {
		r0 = rf(name, k8sSpecFlag, spec)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Environment provides a mock function with given fields: name
func (_m *App) Environment(name string) (*app.EnvironmentSpec, error) {
	ret := _m.Called(name)

	var r0 *app.EnvironmentSpec
	if rf, ok := ret.Get(0).(func(string) *app.EnvironmentSpec); ok {
		r0 = rf(name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*app.EnvironmentSpec)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Environments provides a mock function with given fields:
func (_m *App) Environments() (app.EnvironmentSpecs, error) {
	ret := _m.Called()

	var r0 app.EnvironmentSpecs
	if rf, ok := ret.Get(0).(func() app.EnvironmentSpecs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(app.EnvironmentSpecs)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Fs provides a mock function with given fields:
func (_m *App) Fs() afero.Fs {
	ret := _m.Called()

	var r0 afero.Fs
	if rf, ok := ret.Get(0).(func() afero.Fs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(afero.Fs)
		}
	}

	return r0
}

// Init provides a mock function with given fields:
func (_m *App) Init() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LibPath provides a mock function with given fields: envName
func (_m *App) LibPath(envName string) (string, error) {
	ret := _m.Called(envName)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(envName)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(envName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Libraries provides a mock function with given fields:
func (_m *App) Libraries() app.LibraryRefSpecs {
	ret := _m.Called()

	var r0 app.LibraryRefSpecs
	if rf, ok := ret.Get(0).(func() app.LibraryRefSpecs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(app.LibraryRefSpecs)
		}
	}

	return r0
}

// Registries provides a mock function with given fields:
func (_m *App) Registries() app.RegistryRefSpecs {
	ret := _m.Called()

	var r0 app.RegistryRefSpecs
	if rf, ok := ret.Get(0).(func() app.RegistryRefSpecs); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(app.RegistryRefSpecs)
		}
	}

	return r0
}

// RemoveEnvironment provides a mock function with given fields: name
func (_m *App) RemoveEnvironment(name string) error {
	ret := _m.Called(name)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(name)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RenameEnvironment provides a mock function with given fields: from, to
func (_m *App) RenameEnvironment(from string, to string) error {
	ret := _m.Called(from, to)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(from, to)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Root provides a mock function with given fields:
func (_m *App) Root() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// UpdateTargets provides a mock function with given fields: envName, targets
func (_m *App) UpdateTargets(envName string, targets []string) error {
	ret := _m.Called(envName, targets)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []string) error); ok {
		r0 = rf(envName, targets)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Upgrade provides a mock function with given fields: dryRun
func (_m *App) Upgrade(dryRun bool) error {
	ret := _m.Called(dryRun)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool) error); ok {
		r0 = rf(dryRun)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
