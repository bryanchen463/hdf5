// Copyright ©2017 The Gonum Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hdf5

// #cgo LDFLAGS: -l:libhdf5_debug.a -l:libhdf5_hl_debug.a
// #cgo darwin CFLAGS: -I/usr/local/include
// #cgo darwin LDFLAGS: -L/usr/local/lib
// #cgo linux,!arm64 CFLAGS: -I/usr/local/include, -I/usr/local/HDF_Group/HDF5/2.0.0/include
// #cgo linux,!arm64 LDFLAGS: -L/usr/local/lib, -L/usr/local/HDF_Group/HDF5/2.0.0/lib
// #cgo linux,arm64 CFLAGS: -I/usr/local/include, -I/usr/lib/aarch64-linux-gnu/hdf5/serial/include
// #cgo linux,arm64 LDFLAGS: -L/usr/local/lib, -L/usr/lib/aarch64-linux-gnu/hdf5/serial/
// #include "hdf5.h"
import "C"
