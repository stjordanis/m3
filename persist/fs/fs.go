// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package fs

import (
	"encoding/binary"
)

const (
	infoFileSuffix        = "info"
	indexFileSuffix       = "index"
	summariesFileSuffix   = "summaries"
	bloomFilterFileSuffix = "bloomfilter"
	dataFileSuffix        = "data"
	digestFileSuffix      = "digest"
	checkpointFileSuffix  = "checkpoint"
	filesetFilePrefix     = "fileset"
	commitLogFilePrefix   = "commitlog"
	fileSuffix            = ".db"

	separator                    = "-"
	infoFilePattern              = filesetFilePrefix + separator + "[0-9]*" + separator + infoFileSuffix + fileSuffix
	filesetFilePattern           = filesetFilePrefix + separator + "[0-9]*" + separator + "[a-z]*" + fileSuffix
	commitLogFilePattern         = commitLogFilePrefix + separator + "[0-9]*" + separator + "[0-9]*" + fileSuffix
	commitLogFileForTimeTemplate = commitLogFilePrefix + separator + "%d" + separator + "[0-9]*" + fileSuffix

	// Index ID is int64
	idxLen = 8
)

var (
	// Use an easy marker for out of band analyzing the raw data files
	marker      = []byte{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1}
	markerLen   = len(marker)
	prologueLen = markerLen + idxLen

	// Endianness is little endian
	endianness = binary.LittleEndian
)