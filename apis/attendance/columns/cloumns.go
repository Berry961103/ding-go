// MIT License

// Copyright (c) 2019 Berryhe

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package cloumns 考勤组信息管理 interface
package cloumns

import "github.com/Berry961103/ding"

const (
	apiAttendanceGetSimpleGroups = "/topapi/attendance/getsimplegroups"
)

// AttendanceGetSimpleGroups 批量获取考勤组详情
// See https://ding-doc.dingtalk.com/document#/org-dev-guide/queries-attendance-group-list-details
// POST https://oapi.dingtalk.com/topapi/attendance/getsimplegroups?access_token=ACCESS_TOKEN
func AttendanceGetSimpleGroups(dctx *ding.DingCtx, payload []byte) ([]byte, error) {
	return dctx.HTTPPost(apiAttendanceGetSimpleGroups, payload, ding.DefaultPostDecodeStr)
}