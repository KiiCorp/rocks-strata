//  Copyright (c) 2015, Facebook, Inc.  All rights reserved.
//  This source code is licensed under the BSD-style license found in the
//  LICENSE file in the root directory of this source tree. An additional grant
//  of patent rights can be found in the PATENTS file in the same directory.

package ossstorage

import (
	"testing"

	"github.com/facebookgo/ensure"
	"github.com/naytzyrhc/rocks-strata/strata"
)

func TestOSSStorage(t *testing.T) {
	t.Parallel()
	oss := NewMockOSS(t)
	defer oss.Stop()

	s, err := NewStorageWithMockOSS(oss)
	ensure.Nil(t, err)

	strata.HelpTestStorage(t, s)
}

func TestOSSStorageManyFiles(t *testing.T) {
	t.Parallel()
	oss := NewMockOSS(t)
	defer oss.Stop()

	s, err := NewStorageWithMockOSS(oss)
	ensure.Nil(t, err)

	strata.HelpTestStorageManyFiles(t, s)

}
