package parallels

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/deathly809/gods/queue"
)

// thread data. idk, might need it?
type thread struct {
	myRuntime *runtime
}

func (t *thread) Run() {

	rt := t.myRuntime

	for t.myRuntime.active {

		job := Job(nil)

		rt.mutex.Lock()
		if rt.q.Count() > 0 {
			job = rt.q.Dequeue().(Job)
		}
		rt.mutex.Unlock()

		if job != nil {

			rt.mutex.Lock()
			status := rt.statusMap[job.ID()]
			rt.mutex.Unlock()

			/* Start the job if new */
			if status == Enqueued {
				status = Running
				rt.mutex.Lock()
				rt.statusMap[job.ID()] = status
				rt.mutex.Unlock()
			}

			/* Do one unit of work */
			if status == Running {
				if job.Next() {
					rt.mutex.Lock()
					rt.q.Enqueue(job)
					rt.mutex.Unlock()
				} else {
					rt.mutex.Lock()
					rt.statusMap[job.ID()] = Completed
					rt.mutex.Unlock()
				}
			}
		}
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("Done?")
}

type runtime struct {
	threads   int
	mutex     sync.Mutex
	jID       int64
	q         queue.Queue
	statusMap map[int]JobStatus
	active    bool
}

func (r *runtime) getUUID() int {
	return int(atomic.AddInt64(&r.jID, 1) - 1)
}

func (r *runtime) StartJob(job Job) int {
	result := -1
	if r.active {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		result = r.getUUID()
		job.SetID(result)
		r.statusMap[result] = Enqueued
		r.q.Enqueue(job)
	} else {
		panic("Trying to start a job when runtime is not initialized")
	}
	return result
}

func (r *runtime) StopJob(jobID int) bool {
	result := InvalidJob
	if r.active {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		if t, ok := r.statusMap[jobID]; ok {
			if t == Enqueued || t == Running {
				result, r.statusMap[jobID] = Terminated, Terminated
			}
		}

	}
	return result != InvalidJob
}

// JobStatus returns the status of the job
func (r *runtime) JobStatus(jobID int) JobStatus {
	result := InvalidJob
	if r.active {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		if t, ok := r.statusMap[jobID]; ok {
			result = t
		}
	}
	return result
}

func (r *runtime) Start() {
	if !r.active {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		if !r.active {

			r.q = queue.New()
			r.statusMap = make(map[int]JobStatus)

			for i := 0; i < r.threads; i++ {
				t := &thread{
					myRuntime: r,
				}
				go t.Run()
			}
			r.active = true
		}
	}
}

func (r *runtime) Stop() {
	if r.active {
		r.mutex.Lock()
		defer r.mutex.Unlock()
		r.q = nil
		r.statusMap = nil
		r.active = false
	}

}
