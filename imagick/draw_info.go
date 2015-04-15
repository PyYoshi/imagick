// Copyright 2013 Herbert G. Fischer. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package imagick

/*
#define MAGICKCORE_SIZEOF_DOUBLE_T 12
#define MAGICKCORE_SIZEOF_DOUBLE 12
#include <wand/MagickWand.h>
*/
import "C"

type DrawInfo struct {
	di *C.DrawInfo
}

func NewDrawInfo() *DrawInfo {
	return &DrawInfo{C.AcquireDrawInfo()}
}
