// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	rpcmessages "github.com/digitalbitbox/bitbox-base/middleware/src/rpcmessages"
	mock "github.com/stretchr/testify/mock"
)

// Middleware is an autogenerated mock type for the Middleware type
type Middleware struct {
	mock.Mock
}

// BackupHSMSecret provides a mock function with given fields:
func (_m *Middleware) BackupHSMSecret() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// BackupSysconfig provides a mock function with given fields:
func (_m *Middleware) BackupSysconfig() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// EnableClearnetIBD provides a mock function with given fields: _a0
func (_m *Middleware) EnableClearnetIBD(_a0 rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// EnableRootLogin provides a mock function with given fields: _a0
func (_m *Middleware) EnableRootLogin(_a0 rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// EnableSSHPasswordLogin provides a mock function with given fields: _a0
func (_m *Middleware) EnableSSHPasswordLogin(_a0 rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// EnableTor provides a mock function with given fields: _a0
func (_m *Middleware) EnableTor(_a0 rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// EnableTorElectrs provides a mock function with given fields: _a0
func (_m *Middleware) EnableTorElectrs(_a0 rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// EnableTorMiddleware provides a mock function with given fields: _a0
func (_m *Middleware) EnableTorMiddleware(_a0 rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// EnableTorSSH provides a mock function with given fields: _a0
func (_m *Middleware) EnableTorSSH(_a0 rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.ToggleSettingArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// GetBaseInfo provides a mock function with given fields:
func (_m *Middleware) GetBaseInfo() rpcmessages.GetBaseInfoResponse {
	ret := _m.Called()

	var r0 rpcmessages.GetBaseInfoResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.GetBaseInfoResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.GetBaseInfoResponse)
	}

	return r0
}

// GetBaseUpdateProgress provides a mock function with given fields:
func (_m *Middleware) GetBaseUpdateProgress() rpcmessages.GetBaseUpdateProgressResponse {
	ret := _m.Called()

	var r0 rpcmessages.GetBaseUpdateProgressResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.GetBaseUpdateProgressResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.GetBaseUpdateProgressResponse)
	}

	return r0
}

// GetServiceInfo provides a mock function with given fields:
func (_m *Middleware) GetServiceInfo() rpcmessages.GetServiceInfoResponse {
	ret := _m.Called()

	var r0 rpcmessages.GetServiceInfoResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.GetServiceInfoResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.GetServiceInfoResponse)
	}

	return r0
}

// IsBaseUpdateAvaliable provides a mock function with given fields:
func (_m *Middleware) IsBaseUpdateAvaliable() rpcmessages.IsBaseUpdateAvailableResponse {
	ret := _m.Called()

	var r0 rpcmessages.IsBaseUpdateAvailableResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.IsBaseUpdateAvailableResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.IsBaseUpdateAvailableResponse)
	}

	return r0
}

// RebootBase provides a mock function with given fields:
func (_m *Middleware) RebootBase() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// ReindexBitcoin provides a mock function with given fields:
func (_m *Middleware) ReindexBitcoin() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// RestoreHSMSecret provides a mock function with given fields:
func (_m *Middleware) RestoreHSMSecret() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// RestoreSysconfig provides a mock function with given fields:
func (_m *Middleware) RestoreSysconfig() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// ResyncBitcoin provides a mock function with given fields:
func (_m *Middleware) ResyncBitcoin() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// SetHostname provides a mock function with given fields: _a0
func (_m *Middleware) SetHostname(_a0 rpcmessages.SetHostnameArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.SetHostnameArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// SetLoginPassword provides a mock function with given fields: _a0
func (_m *Middleware) SetLoginPassword(_a0 rpcmessages.SetLoginPasswordArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.SetLoginPasswordArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// SetupStatus provides a mock function with given fields:
func (_m *Middleware) SetupStatus() rpcmessages.SetupStatusResponse {
	ret := _m.Called()

	var r0 rpcmessages.SetupStatusResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.SetupStatusResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.SetupStatusResponse)
	}

	return r0
}

// ShutdownBase provides a mock function with given fields:
func (_m *Middleware) ShutdownBase() rpcmessages.ErrorResponse {
	ret := _m.Called()

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.ErrorResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// SystemEnv provides a mock function with given fields:
func (_m *Middleware) SystemEnv() rpcmessages.GetEnvResponse {
	ret := _m.Called()

	var r0 rpcmessages.GetEnvResponse
	if rf, ok := ret.Get(0).(func() rpcmessages.GetEnvResponse); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(rpcmessages.GetEnvResponse)
	}

	return r0
}

// UpdateBase provides a mock function with given fields: _a0
func (_m *Middleware) UpdateBase(_a0 rpcmessages.UpdateBaseArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.UpdateBaseArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// UserAuthenticate provides a mock function with given fields: _a0
func (_m *Middleware) UserAuthenticate(_a0 rpcmessages.UserAuthenticateArgs) rpcmessages.UserAuthenticateResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.UserAuthenticateResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.UserAuthenticateArgs) rpcmessages.UserAuthenticateResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.UserAuthenticateResponse)
	}

	return r0
}

// UserChangePassword provides a mock function with given fields: _a0
func (_m *Middleware) UserChangePassword(_a0 rpcmessages.UserChangePasswordArgs) rpcmessages.ErrorResponse {
	ret := _m.Called(_a0)

	var r0 rpcmessages.ErrorResponse
	if rf, ok := ret.Get(0).(func(rpcmessages.UserChangePasswordArgs) rpcmessages.ErrorResponse); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(rpcmessages.ErrorResponse)
	}

	return r0
}

// ValidateToken provides a mock function with given fields: token
func (_m *Middleware) ValidateToken(token string) error {
	ret := _m.Called(token)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
