// Copyright 2014 Hu Cong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//
package scheduler

import (
	"visual_spider_go/spider/core/common/request"
)

type Scheduler interface {
	Push(requ *request.Request)
	Poll() *request.Request
	Count() int
}
