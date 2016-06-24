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

func generateFunction(thunk DotThunk, start, end int, mutex *sync.Mutex) func(int) bool {
	hasMore := true
	return func(i int) bool {
		if hasMore {
			thunk.Dot(start, end)
			hasMore = false
			mutex.Lock()
			defer mutex.Unlock()
			thunk.Save()
		}
		return hasMore
	}
}

func generateAndRun(thunkFunc func() DotThunk, length int) {
	_maxThreads := 16
	_minWork := parallels.MinWorkPerThread

	threads := (length + _minWork - 1) / _minWork
	if threads > _maxThreads {
		threads = _maxThreads
		_minWork = (length + threads - 1) / threads
	}

	var lock sync.Mutex
	theJobs := []parallels.Job(nil)

	for threadID := 0; threadID < threads; threadID++ {
		myFunc := func(threadID, start, end int, mutex *sync.Mutex) parallels.Job {
			f := generateFunction(thunkFunc(), start, end, mutex)
			return parallels.CreateIterableJob(fmt.Sprint("Dot Product ", threadID, end), f, 0)
		}
		start := threadID * _minWork
		end := gomath.MinInt(start+_minWork, length)
		theJobs = append(theJobs, myFunc(threadID, start, end, &lock))
	}
	parallels.Run(theJobs)
}

// DotThunk encapsulates common operations among all dot products
//
//	Save	-	When a thread has reached the end of its exection it calls Save
//	Dot		-	On each iteration we call Dot
//
//
type DotThunk struct {
	Save  func()
	Dot   func(int, int)
	Print func(int)
}

// Dot is a generic dot product function.
//
//		We assume that the thunkGen generates a new
//		thunk for each thread to use.  The length is
//		how long the array is
//
func Dot(thunkGen func() DotThunk, length int) {
	generateAndRun(thunkGen, length)
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

		thunk := func() DotThunk {
			localResult := float32(0)
			return DotThunk{
				Save: func() {
					result += localResult
				},
				Dot: func(start, end int) {
					for start < end {
						localResult += a[start] * b[start]
						start++
					}
				},
				Print: func(i int) {
					log.Println(i, ": ", localResult)
				},
			}
		}
		generateAndRun(thunk, length)
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

		thunk := func() DotThunk {
			localResult := float64(0)
			return DotThunk{
				Save: func() {
					result += localResult
				},
				Dot: func(start, end int) {
					for start < end {
						localResult += a[start] * b[start]
						start++
					}
				},
				Print: func(i int) {
					log.Println(i, ": ", localResult)
				},
			}
		}
		generateAndRun(thunk, length)
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

		thunk := func() DotThunk {
			localResult := 0
			return DotThunk{
				Save: func() {
					result += localResult
				},
				Dot: func(start, end int) {
					for start < end {
						localResult += a[start] * b[start]
						start++
					}
				},
				Print: func(i int) {
					log.Println(i, ": ", localResult)
				},
			}
		}
		generateAndRun(thunk, length)
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

		thunk := func() DotThunk {
			localResult := int64(0)
			return DotThunk{
				Save: func() {
					result += localResult
				},
				Dot: func(start, end int) {
					for start < end {
						localResult += int64(a[start]) * int64(b[start])
						start++
					}
				},
				Print: func(i int) {
					log.Println(i, ": ", localResult)
				},
			}
		}
		generateAndRun(thunk, length)
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

		thunk := func() DotThunk {
			localResult := int64(0)
			return DotThunk{
				Save: func() {
					result += localResult
				},
				Dot: func(start, end int) {
					for start < end {
						localResult += a[start] * b[start]
						start++
					}
				},
				Print: func(i int) {
					log.Println(i, ": ", localResult)
				},
			}
		}
		generateAndRun(thunk, length)
	}
	return result
}
