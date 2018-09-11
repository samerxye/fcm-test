package fcm

import "github.com/stretchr/testify/mock"

// httpCMock is an autogenerated mock type for the httpC type
type httpCMock struct {
	mock.Mock
}

// Send provides a mock function with given fields: m
func (_m *httpCMock) Send(m HTTPMessage) (*HTTPResponse, error) {
	ret := _m.Called(m)

	var r0 *HTTPResponse
	if rf, ok := ret.Get(0).(func(HTTPMessage) *HTTPResponse); ok {
		r0 = rf(m)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*HTTPResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(HTTPMessage) error); ok {
		r1 = rf(m)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
