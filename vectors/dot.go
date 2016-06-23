package vectors

import (
	"fmt"
	"log"
	"os"
	"sync"

	"runtime/pprof"

	"github.com/deathly809/gomath"
	"github.com/deathly809/parallels"
)

// This function constructs the appropriate number of jobs to execute.
func wrapper(length int, f func(int, int, *sync.Mutex) parallels.Job) []parallels.Job {

	_maxThreads := 16
	_min := parallels.MinWorkPerThread

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

			theJob := &parallels.IterableJob{
				Name: fmt.Sprint("Dot Product of float32", i),
				TheFunc: func(pos int) bool {
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
		parallels.Run(wrapper(length, f))
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

			theJob := &parallels.IterableJob{
				Name: fmt.Sprint("Dot Product of float32", i),
				TheFunc: func(pos int) bool {
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
		parallels.Run(wrapper(length, f))
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

			theJob := &parallels.IterableJob{
				Name: fmt.Sprint("Dot Product of float32", i),
				TheFunc: func(pos int) bool {
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
		parallels.Run(wrapper(length, f))
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

			theJob := &parallels.IterableJob{
				Name: fmt.Sprint("Dot Product of float32", i),
				TheFunc: func(pos int) bool {
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
		parallels.Run(wrapper(length, f))
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

			theJob := &parallels.IterableJob{
				Name: fmt.Sprint("Dot Product of float32", i),
				TheFunc: func(pos int) bool {
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
		parallels.Run(wrapper(length, f))
	}
	return result
}
