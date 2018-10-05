// Copyright © 2018 NAME HERE <EMAIL ADDRESS>
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

package status

// import (
// 	"errors"
// 	"strings"
// )

const (
	ContainerCreated = "create"

	ContainerStopped = "stop"

	ContainerDied = "die"

	ContainerResize = "resize"

	ContainerAttach = "attach"
)

// const (
// 	// ContainerStatusNone is the zero state of a container; this container has not completed pull
// 	ContainerStatusNone ContainerStatus = iota
// 	// ContainerPulled represents a container which has had the image pulled
// 	ContainerPulled
// 	// ContainerCreated represents a container that has been created
// 	ContainerCreated
// 	// ContainerRunning represents a container that has started
// 	ContainerRunning
// 	// ContainerResourcesProvisioned represents a container that has completed provisioning all of its
// 	// resources. Non-internal containers (containers present in the task definition) transition to
// 	// this state without doing any additional work. However, containers that are added to a task
// 	// by the ECS Agent would possibly need to perform additional actions before they can be
// 	// considered "ready" and contribute to the progress of a task. For example, the "pause" container
// 	// would be provisioned by invoking CNI plugins
// 	ContainerResourcesProvisioned
// 	// ContainerStopped represents a container that has stopped
// 	ContainerStopped
// 	// ContainerZombie is an "impossible" state that is used as the maximum
// 	ContainerZombie
// )

// const (
// 	// ContainerHealthUnknown is the initial status of container health
// 	ContainerHealthUnknown ContainerHealthStatus = iota
// 	// ContainerHealthy represents the status of container health check when returned healthy
// 	ContainerHealthy
// 	// ContainerUnhealthy represents the status of container health check when returned unhealthy
// 	ContainerUnhealthy
// )

// // ContainerStatus is an enumeration of valid states in the container lifecycle
// type ContainerStatus int32

// // ContainerHealthStatus is an enumeration of container health check status
// type ContainerHealthStatus int32

// var containerStatusMap = map[string]ContainerStatus{
// 	"NONE":                  ContainerStatusNone,
// 	"PULLED":                ContainerPulled,
// 	"CREATED":               ContainerCreated,
// 	"RUNNING":               ContainerRunning,
// 	"RESOURCES_PROVISIONED": ContainerResourcesProvisioned,
// 	"STOPPED":               ContainerStopped,
// }

// // BackendStatus returns the container health status recognized by backend
// func (healthStatus ContainerHealthStatus) BackendStatus() string {
// 	switch healthStatus {
// 	case ContainerHealthy:
// 		return "HEALTHY"
// 	case ContainerUnhealthy:
// 		return "UNHEALTHY"
// 	default:
// 		return "UNKNOWN"
// 	}
// }

// // String returns the readable description of the container health status
// func (healthStatus ContainerHealthStatus) String() string {
// 	return healthStatus.BackendStatus()
// }

// // String returns a human readable string representation of this object
// func (cs ContainerStatus) String() string {
// 	for k, v := range containerStatusMap {
// 		if v == cs {
// 			return k
// 		}
// 	}
// 	return "NONE"
// }

// // ShouldReportToBackend returns true if the container status is recognized as a
// // valid state by ECS. Note that not all container statuses are recognized by ECS
// // or map to ECS states
// func (cs *ContainerStatus) ShouldReportToBackend(steadyStateStatus ContainerStatus) bool {
// 	return *cs == steadyStateStatus || *cs == ContainerStopped
// }

// // BackendStatus maps the internal container status in the agent to that in the
// // backend
// func (cs *ContainerStatus) BackendStatus(steadyStateStatus ContainerStatus) ContainerStatus {
// 	if *cs == steadyStateStatus {
// 		return ContainerRunning
// 	}

// 	if *cs == ContainerStopped {
// 		return ContainerStopped
// 	}

// 	return ContainerStatusNone
// }

// // Terminal returns true if the container status is STOPPED
// func (cs ContainerStatus) Terminal() bool {
// 	return cs == ContainerStopped
// }

// // UnmarshalJSON overrides the logic for parsing the JSON-encoded ContainerStatus data
// func (cs *ContainerStatus) UnmarshalJSON(b []byte) error {
// 	if strings.ToLower(string(b)) == "null" {
// 		*cs = ContainerStatusNone
// 		return nil
// 	}
// 	if b[0] != '"' || b[len(b)-1] != '"' {
// 		*cs = ContainerStatusNone
// 		return errors.New("container status unmarshal: status must be a string or null; Got " + string(b))
// 	}
// 	strStatus := string(b[1 : len(b)-1])
// 	// 'UNKNOWN' and 'DEAD' for Compatibility with v1.0.0 state files
// 	if strStatus == "UNKNOWN" {
// 		*cs = ContainerStatusNone
// 		return nil
// 	}
// 	if strStatus == "DEAD" {
// 		*cs = ContainerStopped
// 		return nil
// 	}

// 	stat, ok := containerStatusMap[strStatus]
// 	if !ok {
// 		*cs = ContainerStatusNone
// 		return errors.New("container status unmarshal: unrecognized status")
// 	}
// 	*cs = stat
// 	return nil
// }

// // MarshalJSON overrides the logic for JSON-encoding the ContainerStatus type
// func (cs *ContainerStatus) MarshalJSON() ([]byte, error) {
// 	if cs == nil {
// 		return nil, nil
// 	}
// 	return []byte(`"` + cs.String() + `"`), nil
// }

// // UnmarshalJSON overrides the logic for parsing the JSON-encoded container health data
// func (healthStatus *ContainerHealthStatus) UnmarshalJSON(b []byte) error {
// 	*healthStatus = ContainerHealthUnknown

// 	if strings.ToLower(string(b)) == "null" {
// 		return nil
// 	}
// 	if b[0] != '"' || b[len(b)-1] != '"' {
// 		return errors.New("container health status unmarshal: status must be a string or null; Got " + string(b))
// 	}

// 	strStatus := string(b[1 : len(b)-1])
// 	switch strStatus {
// 	case "UNKNOWN":
// 	// The health status is already set to ContainerHealthUnknown initially
// 	case "HEALTHY":
// 		*healthStatus = ContainerHealthy
// 	case "UNHEALTHY":
// 		*healthStatus = ContainerUnhealthy
// 	default:
// 		return errors.New("container health status unmarshal: unrecognized status: " + string(b))
// 	}
// 	return nil
// }

// // MarshalJSON overrides the logic for JSON-encoding the ContainerHealthStatus type
// func (healthStatus *ContainerHealthStatus) MarshalJSON() ([]byte, error) {
// 	if healthStatus == nil {
// 		return nil, nil
// 	}
// 	return []byte(`"` + healthStatus.String() + `"`), nil
// }

// // IsRunning returns true if the container status is either RUNNING or RESOURCES_PROVISIONED
// func (cs ContainerStatus) IsRunning() bool {
// 	return cs == ContainerRunning || cs == ContainerResourcesProvisioned
// }
