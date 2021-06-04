// Copyright 2021 The gVisor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package strace

import (
	"golang.org/x/sys/unix"
	"gvisor.dev/gvisor/pkg/abi"
)

// ProtectionFlagSet represents the protection to mmap(2).
var ProtectionFlagSet = abi.FlagSet{
	{
		Flag: unix.PROT_READ,
		Name: "PROT_READ",
	},
	{
		Flag: unix.PROT_WRITE,
		Name: "PROT_WRITE",
	},
	{
		Flag: unix.PROT_EXEC,
		Name: "PROT_EXEC",
	},
}

// MmapFlagSet is the set of mmap(2) flags.
var MmapFlagSet = abi.FlagSet{
	{
		Flag: unix.MAP_SHARED,
		Name: "MAP_SHARED",
	},
	{
		Flag: unix.MAP_PRIVATE,
		Name: "MAP_PRIVATE",
	},
	{
		Flag: unix.MAP_FIXED,
		Name: "MAP_FIXED",
	},
	{
		Flag: unix.MAP_ANONYMOUS,
		Name: "MAP_ANONYMOUS",
	},
	{
		Flag: unix.MAP_32BIT,
		Name: "MAP_32BIT",
	},
	{
		Flag: unix.MAP_GROWSDOWN,
		Name: "MAP_GROWSDOWN",
	},
	{
		Flag: unix.MAP_DENYWRITE,
		Name: "MAP_DENYWRITE",
	},
	{
		Flag: unix.MAP_EXECUTABLE,
		Name: "MAP_EXECUTABLE",
	},
	{
		Flag: unix.MAP_LOCKED,
		Name: "MAP_LOCKED",
	},
	{
		Flag: unix.MAP_NORESERVE,
		Name: "MAP_NORESERVE",
	},
	{
		Flag: unix.MAP_POPULATE,
		Name: "MAP_POPULATE",
	},
	{
		Flag: unix.MAP_NONBLOCK,
		Name: "MAP_NONBLOCK",
	},
	{
		Flag: unix.MAP_STACK,
		Name: "MAP_STACK",
	},
	{
		Flag: unix.MAP_HUGETLB,
		Name: "MAP_HUGETLB",
	},
}
