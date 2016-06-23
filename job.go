package parallels

import (
	"sync"
	"time"
)

var (
	MinWorkPerThread = 16 * 1024
)

type IterableJob struct {
	ID      int
	Name    string
	Pos     int
	TheFunc func(int) bool
	lock    sync.Mutex
}

func (j *IterableJob) GetID() int {
	return j.ID
}

func (j *IterableJob) SetID(id int) {
	j.ID = id
}

func (j *IterableJob) GetName() string {
	return j.Name
}

func (j *IterableJob) Next() bool {
	j.lock.Lock()
	curr := j.Pos
	j.Pos++
	j.lock.Unlock()
	return j.TheFunc(curr)
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
