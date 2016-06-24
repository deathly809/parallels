package parallels

import (
	"sync"
	"time"
)

var (
	MinWorkPerThread = 16 * 1024
)

// IterableJob represents the work performed in a
// loop
type IterableJob struct {
	id       int
	idSet    bool
	name     string
	pos      int
	theFunc  func(int) bool
	lock     sync.Mutex
	finished bool
}

// GetID returns the id for this job.
//
//		If the job has been submitted to the runtime then
//		this ID has been set to some internal ID
//
//
//
func (j *IterableJob) GetID() int {
	return j.id
}

// SetID sets the ID if it has not been set previously
//
func (j *IterableJob) SetID(id int) {
	if !j.idSet {
		j.id = id
	}
}

// GetName will return the user given name of this job
func (j *IterableJob) GetName() string {
	return j.name
}

// Next will try to perform the next operation for this
// job
func (j *IterableJob) Next() bool {
	if !j.finished {
		j.lock.Lock()
		curr := j.pos
		j.pos++
		j.lock.Unlock()
		j.finished = !j.theFunc(curr)
	}
	return !j.finished
}

// Run takes in an array of jobs and executes them
func Run(theJobs []Job) {
	rt := GetRuntime()
	rt.Start()
	for i := range theJobs {
		rt.StartJob(theJobs[i])
	}

	for i := range theJobs {
		// Sleep between waiting
		mult := time.Duration(5)
		for rt.JobStatus(theJobs[i].GetID()) != Completed {
			time.Sleep(mult * time.Millisecond)
			if mult < 100 {
				mult += 5
			}
		}
	}
}

// CreateIterableJob returns a new Iterable Job
func CreateIterableJob(name string, f func(i int) bool, currPos int) Job {
	return &IterableJob{
		name:     name,
		theFunc:  f,
		pos:      currPos,
		finished: false,
		idSet:    false,
	}

}
