package parallels

import "sync"

// Job which you want to run in parallel
type Job interface {
	Name() string
	Next() bool
	ID() int
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

// Runtime handles the parallels runtime
type Runtime interface {

	// StartJob starts a job and returns a unique id
	// for the job
	StartJob(theJob Job, async bool) int
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
			globalRuntime = &runtime{threads: 25}
		}
		mutex.Unlock()
	}
	return globalRuntime
}
