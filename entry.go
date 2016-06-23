package parallels

import (
	rt "runtime"
	"sync"
)

// Job which you want to run in parallel
type Job interface {
	GetName() string
	Next() bool
	GetID() int
	SetID(int)
}

// JobStatus represents the current status of a
// job
type JobStatus int

/*
 * The different statuses of a Job
 */
const (
	InvalidJob = JobStatus(iota) // Could not find job
	Enqueued   = JobStatus(iota) // In the queue, waiting to run
	Running    = JobStatus(iota) // Running
	Terminated = JobStatus(iota) // Error/forced quit
	Completed  = JobStatus(iota) // Completed successfully
)

// StatusByName returns a status value by its string representation
func StatusByName(status JobStatus) string {
	result := "Invalid Status"
	switch status {
	case InvalidJob:
		result = "InvalidJob"
	case Enqueued:
		result = "Enqueued"
	case Running:
		result = "Running"
	case Terminated:
		result = "Terminated"
	case Completed:
		result = "Completed"
	}
	return result
}

// Runtime handles the parallels runtime
type Runtime interface {

	// StartJob starts a job and returns a UID
	// for the job
	StartJob(job Job) int

	// StopJob stops the given job
	StopJob(jobID int) bool

	// JobStatus returns the status of the job
	JobStatus(jobID int) JobStatus

	// Start the parallels runtime
	Start()

	// Stop the parallels runtime
	Stop()
}

var (
	globalRuntime Runtime
	mutex         sync.Mutex
)

// GetRuntime returns the current runtime
func GetRuntime() Runtime {
	if globalRuntime == nil {
		mutex.Lock()
		if globalRuntime == nil {
			rt.GOMAXPROCS(rt.NumCPU())
			globalRuntime = &runtime{threads: 10}
		}
		mutex.Unlock()
	}
	return globalRuntime
}
