package vectors

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"runtime/pprof"

	"github.com/deathly809/gomath"
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

func wrapper(length int, f func(int, int, *sync.Mutex) parallels.Job) []parallels.Job {
	_maxThreads := 16
	_min := _MinWorkPerThread

	threads := (length + _min - 1) / _min

	if threads > _maxThreads {
		threads = _maxThreads
		_min = (length + threads - 1) / threads
	}

	var lock sync.Mutex

	theJobs := []parallels.Job(nil)

	for i := 0; i < threads; i++ {
		theJobs = append(theJobs, f(i, _min, &lock))
	}

	return theJobs
}

// DotFloat32 performs a dot product on two arrays of float32
func DotFloat32(a, b []float32) float32 {

	f, err := os.Create("profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	result := float32(0.0)

	if a != nil && b != nil {

		length := gomath.MinInt(len(a), len(b))

		a = a[:length]
		b = b[:length]

		f := func(i, _min int, lock *sync.Mutex) parallels.Job {
			tLen := gomath.MinInt(_min, len(a))

			aData := a[:tLen]
			bData := b[:tLen]

			a = a[tLen:]
			b = b[tLen:]

			theJob := &job{
				name: fmt.Sprint("Dot Product of float32", i),
				theFunc: func(pos int) bool {
					res := false
					r := float32(0.0)
					pos = pos * _min

					if pos < tLen {
						aArr := aData[pos:]
						bArr := bData[pos:]
						l := gomath.MinInt(_min, tLen-pos)
						if l > 0 {
							for i := 0; i < l; i++ {
								r += aArr[i] * bArr[i]
							}
							lock.Lock()
							result += r
							lock.Unlock()
							res = true
						}
					}
					return res
				},
			}
			return theJob
		}
		run(wrapper(length, f))
	} else {
		panic(fmt.Sprint("Invalid parameters:", a, b))
	}
	return result
}

// DotFloat64 performs a dot product on two arrays of float64
func DotFloat64(a, b []float64) float64 {
	result := 0.0
	if a != nil && b != nil {

		length := gomath.MinInt(len(a), len(b))

		a = a[:length]
		b = b[:length]

		f := func(i, _min int, lock *sync.Mutex) parallels.Job {
			tLen := gomath.MinInt(_min, len(a))

			aData := a[:tLen]
			bData := b[:tLen]

			a = a[tLen:]
			b = b[tLen:]

			theJob := &job{
				name: fmt.Sprint("Dot Product of float32", i),
				theFunc: func(pos int) bool {
					res := false
					r := float64(0.0)
					pos = pos * _min

					if pos < tLen {
						aArr := aData[pos:]
						bArr := bData[pos:]
						l := gomath.MinInt(_min, tLen-pos)
						if l > 0 {
							for i := 0; i < l; i++ {
								r += aArr[i] * bArr[i]
							}
							lock.Lock()
							result += r
							lock.Unlock()
							res = true
						}
					}
					return res
				},
			}
			return theJob
		}
		run(wrapper(length, f))
	}
	return result
}

// DotInt performs a dot product on two arrays of float64
func DotInt(a, b []int) int {
	result := 0
	if a != nil && b != nil {

		length := gomath.MinInt(len(a), len(b))

		a = a[:length]
		b = b[:length]

		f := func(i, _min int, lock *sync.Mutex) parallels.Job {
			tLen := gomath.MinInt(_min, len(a))

			aData := a[:tLen]
			bData := b[:tLen]

			a = a[tLen:]
			b = b[tLen:]

			theJob := &job{
				name: fmt.Sprint("Dot Product of float32", i),
				theFunc: func(pos int) bool {
					res := false
					r := int(0.0)
					pos = pos * _min

					if pos < tLen {
						aArr := aData[pos:]
						bArr := bData[pos:]
						l := gomath.MinInt(_min, tLen-pos)
						if l > 0 {
							for i := 0; i < l; i++ {
								r += aArr[i] * bArr[i]
							}
							lock.Lock()
							result += r
							lock.Unlock()
							res = true
						}
					}
					return res
				},
			}
			return theJob
		}
		run(wrapper(length, f))
	}
	return result
}

// DotInt32 performs a dot product on two arrays of float64
func DotInt32(a, b []int32) int32 {
	result := int64(0)

	if a != nil && b != nil {
		length := gomath.MinInt(len(a), len(b))

		a = a[:length]
		b = b[:length]

		f := func(i, _min int, lock *sync.Mutex) parallels.Job {
			tLen := gomath.MinInt(_min, len(a))

			aData := a[:tLen]
			bData := b[:tLen]

			a = a[tLen:]
			b = b[tLen:]

			theJob := &job{
				name: fmt.Sprint("Dot Product of float32", i),
				theFunc: func(pos int) bool {
					res := false
					r := int64(0)
					pos = pos * _min

					if pos < tLen {
						aArr := aData[pos:]
						bArr := bData[pos:]
						l := gomath.MinInt(_min, tLen-pos)
						if l > 0 {
							for i := 0; i < l; i++ {
								r += int64(aArr[i] * bArr[i])
							}
							lock.Lock()
							result += r
							lock.Unlock()
							res = true
						}
					}
					return res
				},
			}
			return theJob
		}
		run(wrapper(length, f))
	}
	return int32(result)
}

// DotInt64 performs a dot product on two arrays of float64
func DotInt64(a, b []int64) int64 {
	result := int64(0)
	if a != nil && b != nil {
		length := gomath.MinInt(len(a), len(b))

		a = a[:length]
		b = b[:length]

		f := func(i, _min int, lock *sync.Mutex) parallels.Job {
			tLen := gomath.MinInt(_min, len(a))

			aData := a[:tLen]
			bData := b[:tLen]

			a = a[tLen:]
			b = b[tLen:]

			theJob := &job{
				name: fmt.Sprint("Dot Product of float32", i),
				theFunc: func(pos int) bool {
					res := false
					r := int64(0.0)
					pos = pos * _min

					if pos < tLen {
						aArr := aData[pos:]
						bArr := bData[pos:]
						l := gomath.MinInt(_min, tLen-pos)
						if l > 0 {
							for i := 0; i < l; i++ {
								r += aArr[i] * bArr[i]
							}
							lock.Lock()
							result += r
							lock.Unlock()
							res = true
						}
					}
					return res
				},
			}
			return theJob
		}
		run(wrapper(length, f))
	}
	return result
}
