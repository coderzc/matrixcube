// Copyright 2021 Matrix Origin
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import fz "github.com/matrixorigin/matrixcube/chaostesting"

func init() {
	fz.RegisterAction(ActionSet{})
	fz.RegisterAction(ActionGet{})
	fz.RegisterAction(ActionStopNode{})
	fz.RegisterAction(ActionRestartNode{})
	fz.RegisterAction(ActionCrashNode{})
	fz.RegisterAction(ActionIsolateNode{})
	fz.RegisterAction(ActionFullyIsolateNode{})
}

type ActionSet struct {
	ClientID int `xml:",attr"`
	Key      int `xml:",attr"`
	Value    int `xml:",attr"`
}

type ActionGet struct {
	ClientID int `xml:",attr"`
	Key      int `xml:",attr"`
}

type ActionStopNode struct {
	NodeID fz.NodeID `xml:",attr"`
}

type ActionRestartNode struct {
	NodeID fz.NodeID `xml:",attr"`
}

type ActionCrashNode struct {
	NodeID fz.NodeID `xml:",attr"`
}

type ActionIsolateNode struct {
	NodeID  fz.NodeID   `xml:",attr"`
	Between []fz.NodeID `xml:"Between"`
}

type ActionFullyIsolateNode struct {
	NodeID fz.NodeID `xml:",attr"`
}
