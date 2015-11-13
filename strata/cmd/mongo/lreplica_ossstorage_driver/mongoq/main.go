//  Copyright (c) 2015, Facebook, Inc.  All rights reserved.
//  This source code is licensed under the BSD-style license found in the
//  LICENSE file in the root directory of this source tree. An additional grant
//  of patent rights can be found in the PATENTS file in the same directory.

package main

import (
	"github.com/facebookgo/rocks-strata/strata/mongo"
	"github.com/naytzyrhc/rocks-strata/strata/cmd/mongo/lreplica_ossstorage_driver/lrossdriver"
)

func main() {
	mongoq.RunCLI(lrossdriver.DriverFactory{Ops: &lrossdriver.Options{}})
}
