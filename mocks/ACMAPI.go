// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import acm "github.com/aws/aws-sdk-go/service/acm"

import aws "github.com/aws/aws-sdk-go/aws"
import mock "github.com/stretchr/testify/mock"
import request "github.com/aws/aws-sdk-go/aws/request"

// ACMAPI is an autogenerated mock type for the ACMAPI type
type ACMAPI struct {
	mock.Mock
}

// AddTagsToCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) AddTagsToCertificate(_a0 *acm.AddTagsToCertificateInput) (*acm.AddTagsToCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.AddTagsToCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.AddTagsToCertificateInput) *acm.AddTagsToCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.AddTagsToCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.AddTagsToCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AddTagsToCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) AddTagsToCertificateRequest(_a0 *acm.AddTagsToCertificateInput) (*request.Request, *acm.AddTagsToCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.AddTagsToCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.AddTagsToCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.AddTagsToCertificateInput) *acm.AddTagsToCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.AddTagsToCertificateOutput)
		}
	}

	return r0, r1
}

// AddTagsToCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) AddTagsToCertificateWithContext(_a0 aws.Context, _a1 *acm.AddTagsToCertificateInput, _a2 ...request.Option) (*acm.AddTagsToCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.AddTagsToCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.AddTagsToCertificateInput, ...request.Option) *acm.AddTagsToCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.AddTagsToCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.AddTagsToCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) DeleteCertificate(_a0 *acm.DeleteCertificateInput) (*acm.DeleteCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.DeleteCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.DeleteCertificateInput) *acm.DeleteCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.DeleteCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.DeleteCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DeleteCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) DeleteCertificateRequest(_a0 *acm.DeleteCertificateInput) (*request.Request, *acm.DeleteCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.DeleteCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.DeleteCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.DeleteCertificateInput) *acm.DeleteCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.DeleteCertificateOutput)
		}
	}

	return r0, r1
}

// DeleteCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) DeleteCertificateWithContext(_a0 aws.Context, _a1 *acm.DeleteCertificateInput, _a2 ...request.Option) (*acm.DeleteCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.DeleteCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.DeleteCertificateInput, ...request.Option) *acm.DeleteCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.DeleteCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.DeleteCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) DescribeCertificate(_a0 *acm.DescribeCertificateInput) (*acm.DescribeCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.DescribeCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.DescribeCertificateInput) *acm.DescribeCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.DescribeCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.DescribeCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// DescribeCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) DescribeCertificateRequest(_a0 *acm.DescribeCertificateInput) (*request.Request, *acm.DescribeCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.DescribeCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.DescribeCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.DescribeCertificateInput) *acm.DescribeCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.DescribeCertificateOutput)
		}
	}

	return r0, r1
}

// DescribeCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) DescribeCertificateWithContext(_a0 aws.Context, _a1 *acm.DescribeCertificateInput, _a2 ...request.Option) (*acm.DescribeCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.DescribeCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.DescribeCertificateInput, ...request.Option) *acm.DescribeCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.DescribeCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.DescribeCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) ExportCertificate(_a0 *acm.ExportCertificateInput) (*acm.ExportCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.ExportCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.ExportCertificateInput) *acm.ExportCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ExportCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.ExportCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ExportCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) ExportCertificateRequest(_a0 *acm.ExportCertificateInput) (*request.Request, *acm.ExportCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.ExportCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.ExportCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.ExportCertificateInput) *acm.ExportCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.ExportCertificateOutput)
		}
	}

	return r0, r1
}

// ExportCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) ExportCertificateWithContext(_a0 aws.Context, _a1 *acm.ExportCertificateInput, _a2 ...request.Option) (*acm.ExportCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.ExportCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.ExportCertificateInput, ...request.Option) *acm.ExportCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ExportCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.ExportCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) GetCertificate(_a0 *acm.GetCertificateInput) (*acm.GetCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.GetCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.GetCertificateInput) *acm.GetCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.GetCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.GetCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) GetCertificateRequest(_a0 *acm.GetCertificateInput) (*request.Request, *acm.GetCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.GetCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.GetCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.GetCertificateInput) *acm.GetCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.GetCertificateOutput)
		}
	}

	return r0, r1
}

// GetCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) GetCertificateWithContext(_a0 aws.Context, _a1 *acm.GetCertificateInput, _a2 ...request.Option) (*acm.GetCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.GetCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.GetCertificateInput, ...request.Option) *acm.GetCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.GetCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.GetCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) ImportCertificate(_a0 *acm.ImportCertificateInput) (*acm.ImportCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.ImportCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.ImportCertificateInput) *acm.ImportCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ImportCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.ImportCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ImportCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) ImportCertificateRequest(_a0 *acm.ImportCertificateInput) (*request.Request, *acm.ImportCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.ImportCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.ImportCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.ImportCertificateInput) *acm.ImportCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.ImportCertificateOutput)
		}
	}

	return r0, r1
}

// ImportCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) ImportCertificateWithContext(_a0 aws.Context, _a1 *acm.ImportCertificateInput, _a2 ...request.Option) (*acm.ImportCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.ImportCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.ImportCertificateInput, ...request.Option) *acm.ImportCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ImportCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.ImportCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCertificates provides a mock function with given fields: _a0
func (_m *ACMAPI) ListCertificates(_a0 *acm.ListCertificatesInput) (*acm.ListCertificatesOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.ListCertificatesOutput
	if rf, ok := ret.Get(0).(func(*acm.ListCertificatesInput) *acm.ListCertificatesOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ListCertificatesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.ListCertificatesInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListCertificatesPages provides a mock function with given fields: _a0, _a1
func (_m *ACMAPI) ListCertificatesPages(_a0 *acm.ListCertificatesInput, _a1 func(*acm.ListCertificatesOutput, bool) bool) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func(*acm.ListCertificatesInput, func(*acm.ListCertificatesOutput, bool) bool) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListCertificatesPagesWithContext provides a mock function with given fields: _a0, _a1, _a2, _a3
func (_m *ACMAPI) ListCertificatesPagesWithContext(_a0 aws.Context, _a1 *acm.ListCertificatesInput, _a2 func(*acm.ListCertificatesOutput, bool) bool, _a3 ...request.Option) error {
	_va := make([]interface{}, len(_a3))
	for _i := range _a3 {
		_va[_i] = _a3[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1, _a2)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.ListCertificatesInput, func(*acm.ListCertificatesOutput, bool) bool, ...request.Option) error); ok {
		r0 = rf(_a0, _a1, _a2, _a3...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ListCertificatesRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) ListCertificatesRequest(_a0 *acm.ListCertificatesInput) (*request.Request, *acm.ListCertificatesOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.ListCertificatesInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.ListCertificatesOutput
	if rf, ok := ret.Get(1).(func(*acm.ListCertificatesInput) *acm.ListCertificatesOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.ListCertificatesOutput)
		}
	}

	return r0, r1
}

// ListCertificatesWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) ListCertificatesWithContext(_a0 aws.Context, _a1 *acm.ListCertificatesInput, _a2 ...request.Option) (*acm.ListCertificatesOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.ListCertificatesOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.ListCertificatesInput, ...request.Option) *acm.ListCertificatesOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ListCertificatesOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.ListCertificatesInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) ListTagsForCertificate(_a0 *acm.ListTagsForCertificateInput) (*acm.ListTagsForCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.ListTagsForCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.ListTagsForCertificateInput) *acm.ListTagsForCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ListTagsForCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.ListTagsForCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListTagsForCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) ListTagsForCertificateRequest(_a0 *acm.ListTagsForCertificateInput) (*request.Request, *acm.ListTagsForCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.ListTagsForCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.ListTagsForCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.ListTagsForCertificateInput) *acm.ListTagsForCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.ListTagsForCertificateOutput)
		}
	}

	return r0, r1
}

// ListTagsForCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) ListTagsForCertificateWithContext(_a0 aws.Context, _a1 *acm.ListTagsForCertificateInput, _a2 ...request.Option) (*acm.ListTagsForCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.ListTagsForCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.ListTagsForCertificateInput, ...request.Option) *acm.ListTagsForCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ListTagsForCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.ListTagsForCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTagsFromCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) RemoveTagsFromCertificate(_a0 *acm.RemoveTagsFromCertificateInput) (*acm.RemoveTagsFromCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.RemoveTagsFromCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.RemoveTagsFromCertificateInput) *acm.RemoveTagsFromCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.RemoveTagsFromCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.RemoveTagsFromCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RemoveTagsFromCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) RemoveTagsFromCertificateRequest(_a0 *acm.RemoveTagsFromCertificateInput) (*request.Request, *acm.RemoveTagsFromCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.RemoveTagsFromCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.RemoveTagsFromCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.RemoveTagsFromCertificateInput) *acm.RemoveTagsFromCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.RemoveTagsFromCertificateOutput)
		}
	}

	return r0, r1
}

// RemoveTagsFromCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) RemoveTagsFromCertificateWithContext(_a0 aws.Context, _a1 *acm.RemoveTagsFromCertificateInput, _a2 ...request.Option) (*acm.RemoveTagsFromCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.RemoveTagsFromCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.RemoveTagsFromCertificateInput, ...request.Option) *acm.RemoveTagsFromCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.RemoveTagsFromCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.RemoveTagsFromCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestCertificate provides a mock function with given fields: _a0
func (_m *ACMAPI) RequestCertificate(_a0 *acm.RequestCertificateInput) (*acm.RequestCertificateOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.RequestCertificateOutput
	if rf, ok := ret.Get(0).(func(*acm.RequestCertificateInput) *acm.RequestCertificateOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.RequestCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.RequestCertificateInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestCertificateRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) RequestCertificateRequest(_a0 *acm.RequestCertificateInput) (*request.Request, *acm.RequestCertificateOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.RequestCertificateInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.RequestCertificateOutput
	if rf, ok := ret.Get(1).(func(*acm.RequestCertificateInput) *acm.RequestCertificateOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.RequestCertificateOutput)
		}
	}

	return r0, r1
}

// RequestCertificateWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) RequestCertificateWithContext(_a0 aws.Context, _a1 *acm.RequestCertificateInput, _a2 ...request.Option) (*acm.RequestCertificateOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.RequestCertificateOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.RequestCertificateInput, ...request.Option) *acm.RequestCertificateOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.RequestCertificateOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.RequestCertificateInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResendValidationEmail provides a mock function with given fields: _a0
func (_m *ACMAPI) ResendValidationEmail(_a0 *acm.ResendValidationEmailInput) (*acm.ResendValidationEmailOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.ResendValidationEmailOutput
	if rf, ok := ret.Get(0).(func(*acm.ResendValidationEmailInput) *acm.ResendValidationEmailOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ResendValidationEmailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.ResendValidationEmailInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ResendValidationEmailRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) ResendValidationEmailRequest(_a0 *acm.ResendValidationEmailInput) (*request.Request, *acm.ResendValidationEmailOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.ResendValidationEmailInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.ResendValidationEmailOutput
	if rf, ok := ret.Get(1).(func(*acm.ResendValidationEmailInput) *acm.ResendValidationEmailOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.ResendValidationEmailOutput)
		}
	}

	return r0, r1
}

// ResendValidationEmailWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) ResendValidationEmailWithContext(_a0 aws.Context, _a1 *acm.ResendValidationEmailInput, _a2 ...request.Option) (*acm.ResendValidationEmailOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.ResendValidationEmailOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.ResendValidationEmailInput, ...request.Option) *acm.ResendValidationEmailOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.ResendValidationEmailOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.ResendValidationEmailInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCertificateOptions provides a mock function with given fields: _a0
func (_m *ACMAPI) UpdateCertificateOptions(_a0 *acm.UpdateCertificateOptionsInput) (*acm.UpdateCertificateOptionsOutput, error) {
	ret := _m.Called(_a0)

	var r0 *acm.UpdateCertificateOptionsOutput
	if rf, ok := ret.Get(0).(func(*acm.UpdateCertificateOptionsInput) *acm.UpdateCertificateOptionsOutput); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.UpdateCertificateOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*acm.UpdateCertificateOptionsInput) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateCertificateOptionsRequest provides a mock function with given fields: _a0
func (_m *ACMAPI) UpdateCertificateOptionsRequest(_a0 *acm.UpdateCertificateOptionsInput) (*request.Request, *acm.UpdateCertificateOptionsOutput) {
	ret := _m.Called(_a0)

	var r0 *request.Request
	if rf, ok := ret.Get(0).(func(*acm.UpdateCertificateOptionsInput) *request.Request); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*request.Request)
		}
	}

	var r1 *acm.UpdateCertificateOptionsOutput
	if rf, ok := ret.Get(1).(func(*acm.UpdateCertificateOptionsInput) *acm.UpdateCertificateOptionsOutput); ok {
		r1 = rf(_a0)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*acm.UpdateCertificateOptionsOutput)
		}
	}

	return r0, r1
}

// UpdateCertificateOptionsWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) UpdateCertificateOptionsWithContext(_a0 aws.Context, _a1 *acm.UpdateCertificateOptionsInput, _a2 ...request.Option) (*acm.UpdateCertificateOptionsOutput, error) {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *acm.UpdateCertificateOptionsOutput
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.UpdateCertificateOptionsInput, ...request.Option) *acm.UpdateCertificateOptionsOutput); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*acm.UpdateCertificateOptionsOutput)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(aws.Context, *acm.UpdateCertificateOptionsInput, ...request.Option) error); ok {
		r1 = rf(_a0, _a1, _a2...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WaitUntilCertificateValidated provides a mock function with given fields: _a0
func (_m *ACMAPI) WaitUntilCertificateValidated(_a0 *acm.DescribeCertificateInput) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*acm.DescribeCertificateInput) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// WaitUntilCertificateValidatedWithContext provides a mock function with given fields: _a0, _a1, _a2
func (_m *ACMAPI) WaitUntilCertificateValidatedWithContext(_a0 aws.Context, _a1 *acm.DescribeCertificateInput, _a2 ...request.WaiterOption) error {
	_va := make([]interface{}, len(_a2))
	for _i := range _a2 {
		_va[_i] = _a2[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _a0, _a1)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(aws.Context, *acm.DescribeCertificateInput, ...request.WaiterOption) error); ok {
		r0 = rf(_a0, _a1, _a2...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
