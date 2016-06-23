package parallels

import (
	"sync"
	"time"

	"github.com/deathly809/parallels"
)

var (
	_MinWorkPerThread = 16 * 1024
)

type job struct {
	id      int
	name    string
	pos     int
	theFunc func(int) bool
	lock    sync.Mutex
}

func (j *job) ID() int {
	return j.id
}

func (j *job) SetID(id int) {
	j.id = id
}

func (j *job) Name() string {
	return j.name
}

func (j *job) Next() bool {
	j.lock.Lock()
	curr := j.pos
	j.pos++
	j.lock.Unlock()
	return j.theFunc(curr)
}

func run(theJobs []parallels.Job) {
	rt := parallels.GetRuntime()
	rt.Start()
	for i := range theJobs {
		rt.StartJob(theJobs[i])
	}

	for i := range theJobs {
		// Sleep between waiting
		mult := time.Duration(5)
		for rt.JobStatus(theJobs[i].ID()) != parallels.Completed {
			time.Sleep(mult * time.Millisecond)
			if mult < 100 {
				mult += 5
			}
		}
	}
}
