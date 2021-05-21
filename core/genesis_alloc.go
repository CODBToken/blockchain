// Copyright 2017 The JINBAO Authors
// This file is part of the JINBAO library.
//
// The JINBAO library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The JINBAO library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the JINBAO library. If not, see <http://www.gnu.org/licenses/>.

package core

// Constants containing the genesis allocation of built-in genesis blocks.
// Their content is an RLP-encoded list of (address, balance) tuples.
// Use mkalloc.go to create/update them.

// nolint: misspell
const mainnetAllocData = "\xe2\xe1\x94G\x88\xa3\x9fe\xb7\xa8\xe1\t\xde,/\x81\xd7kBPcx\x12\x8bR\xb7\xd3g\x8f/\xd7m\xe8\x00\x00"

const testnetAllocData = "\xe2\xe1\x94G\x88\xa3\x9fe\xb7\xa8\xe1\t\xde,/\x81\xd7kBPcx\x12\x8bR\xb7\xd3g\x8f/\xd7m\xe8\x00\x00"

const devnetAllocData = "\xe2\xe1\x94G\x88\xa3\x9fe\xb7\xa8\xe1\t\xde,/\x81\xd7kBPcx\x12\x8bR\xb7\xd3g\x8f/\xd7m\xe8\x00\x00"
