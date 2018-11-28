// Autogenerated by Thrift Compiler (0.12.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package aurora

import (
	"bytes"
	"context"
	"fmt"
	"reflect"

	"git.apache.org/thrift.git/lib/go/thrift"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = context.Background
var _ = reflect.DeepEqual
var _ = bytes.Equal

const AURORA_EXECUTOR_NAME = "AuroraExecutor"

var ACTIVE_STATES []ScheduleStatus
var SLAVE_ASSIGNED_STATES []ScheduleStatus
var LIVE_STATES []ScheduleStatus
var TERMINAL_STATES []ScheduleStatus

const GOOD_IDENTIFIER_PATTERN = "^[\\w\\-\\.]+$"
const GOOD_IDENTIFIER_PATTERN_JVM = "^[\\w\\-\\.]+$"
const GOOD_IDENTIFIER_PATTERN_PYTHON = "^[\\w\\-\\.]+$"

var ACTIVE_JOB_UPDATE_STATES []JobUpdateStatus
var AWAITNG_PULSE_JOB_UPDATE_STATES []JobUpdateStatus

const BYPASS_LEADER_REDIRECT_HEADER_NAME = "Bypass-Leader-Redirect"
const TASK_FILESYSTEM_MOUNT_POINT = "taskfs"

func init() {
	ACTIVE_STATES = []ScheduleStatus{
		9, 17, 6, 0, 13, 12, 2, 1, 18, 16}

	SLAVE_ASSIGNED_STATES = []ScheduleStatus{
		9, 17, 6, 13, 12, 2, 18, 1}

	LIVE_STATES = []ScheduleStatus{
		6, 13, 12, 17, 18, 2}

	TERMINAL_STATES = []ScheduleStatus{
		4, 3, 5, 7}

	ACTIVE_JOB_UPDATE_STATES = []JobUpdateStatus{
		0, 1, 2, 3, 9, 10}

	AWAITNG_PULSE_JOB_UPDATE_STATES = []JobUpdateStatus{
		9, 10}

}
