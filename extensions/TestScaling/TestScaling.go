/* rpi3.go: extension adds special fields for Bitscope/Raspberry Pi 3B(+) management
 *
 * Author: J. Lowell Wofford <lowell@lanl.gov>
 *
 * This software is open source software available under the BSD-3 license.
 * Copyright (c) 2018, Triad National Security, LLC
 * See LICENSE file for details.
 */

package vbox

import (
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	"github.com/hpc/kraken/core"
	pb "github.com/hpc/kraken/extensions/TestScaling/proto"
	"github.com/hpc/kraken/lib"
)

//go:generate protoc -I ../../core/proto/include -I proto --go_out=plugins=grpc:proto proto/Test.proto

/////////////////
// TestScaling Object /
///////////////

var _ lib.Extension = TestScaling{}

type TestScaling struct{}

func (TestScaling) New() proto.Message {
	return &pb.TestScaling{}
}

func (r TestScaling) Name() string {
	a, _ := ptypes.MarshalAny(r.New())
	return a.GetTypeUrl()
}

func init() {
	core.Registry.RegisterExtension(TestScaling{})
}