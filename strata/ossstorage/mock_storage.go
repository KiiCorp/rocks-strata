//  Copyright (c) 2015, Facebook, Inc.  All rights reserved.
//  This source code is licensed under the BSD-style license found in the
//  LICENSE file in the root directory of this source tree. An additional grant
//  of patent rights can be found in the PATENTS file in the same directory.

package ossstorage

import (
	"testing"

	"github.com/PinIdea/oss-aliyun-go"
	"github.com/facebookgo/ensure"
)

// MockOSS mocks OSSStorage
type MockOSS struct {
	auth   oss.Auth
	region oss.Region
	srv    *osstest.Server
	config *osstest.Config
}

// Start starts server used for MockS3
func (s *MockOSS) Start(t *testing.T) {
	srv, err := osstest.NewServer(s.config)
	ensure.Nil(t, err)
	ensure.NotNil(t, srv)

	s.srv = srv
	s.region = oss.Region{
		Name:                 "faux-region-1",
		OSSEndpoint:           srv.URL(),
		OSSLocationConstraint: true, // s3test server requires a LocationConstraint
	}
}

// Stop stops server used for MockOSS
func (s *MockOSS) Stop() {
	s.srv.Quit()
}

// NewMockS3 constructs a MockS3
func NewMockOSS(t *testing.T) *MockOSS {
	m := MockOSS{}
	m.Start(t)
	return &m
}

// NewStorageWithMockS3 constructs an S3Storage that uses MockS3
func NewStorageWithMockOSS(s *MockOSS) (*OSSStorage, error) {
	return NewOSSStorage(s.region, s.auth, "testbucket", "test", s3.Private)
}
