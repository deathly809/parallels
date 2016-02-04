package vectors

import (
	"sync"

	"github.com/deathly809/gomath"
	"github.com/deathly809/parallels"
)

var (
	_MinWorkPerThread = 1000
)

type job struct {
	id      int
	name    string
	pos     int
	theFunc func(int) bool
}

func (j *job) ID() int {
	return j.id
}

func (j *job) Name() string {
	return j.name
}

func (j *job) Next() bool {
	var lock sync.Mutex
	lock.Lock()
	curr := j.pos
	j.pos++
	lock.Unlock()
	return j.theFunc(curr)
}

// Code clones ahead...

// DotFloat32 performs a dot product on two arrays of float32
func DotFloat32(a, b []float32) float32 {
	result := float32(0.0)

	if a != nil && b != nil {
		l := gomath.MinInt(len(a), len(b))

		a = a[:l]
		b = b[:l]

		var lock sync.Mutex

		theJob := &job{
			name: "Dot Product of float32",
			theFunc: func(pos int) bool {
				res := false
				r := float32(0.0)
				pos = pos * _MinWorkPerThread

				if pos < len(a) && pos < len(b) {
					aArr := a[pos:]
					bArr := b[pos:]
					l := gomath.MinInt(_MinWorkPerThread, gomath.MinInt(len(aArr), len(bArr)))
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
		parallels.GetRuntime().Start()
		parallels.GetRuntime().StartJob(theJob, false)
	}
	return result
}

// DotFloat64 performs a dot product on two arrays of float64
func DotFloat64(a, b []float64) float64 {
	result := 0.0
	if a != nil && b != nil {
		l := gomath.MinInt(len(a), len(b))

		a = a[:l]
		b = b[:l]

		var lock sync.Mutex

		theJob := &job{
			name: "Dot Product of float32",
			theFunc: func(pos int) bool {
				res := false
				r := 0.0
				pos = pos * _MinWorkPerThread

				if pos < len(a) && pos < len(b) {
					aArr := a[pos:]
					bArr := b[pos:]
					l := gomath.MinInt(_MinWorkPerThread, gomath.MinInt(len(aArr), len(bArr)))
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

		parallels.GetRuntime().StartJob(theJob, false)
	}
	return result
}

// DotInt performs a dot product on two arrays of float64
func DotInt(a, b []int) int {
	result := 0
	if a != nil && b != nil {
		l := gomath.MinInt(len(a), len(b))

		a = a[:l]
		b = b[:l]

		var lock sync.Mutex

		theJob := &job{
			name: "Dot Product of float32",
			theFunc: func(pos int) bool {
				res := false
				r := 0
				pos = pos * _MinWorkPerThread

				if pos < len(a) && pos < len(b) {
					aArr := a[pos:]
					bArr := b[pos:]
					l := gomath.MinInt(_MinWorkPerThread, gomath.MinInt(len(aArr), len(bArr)))
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
		parallels.GetRuntime().StartJob(theJob, false)
	}
	return result
}

// DotInt32 performs a dot product on two arrays of float64
func DotInt32(a, b []int32) int32 {
	result := int32(0)
	if a != nil && b != nil {
		l := gomath.MinInt(len(a), len(b))

		a = a[:l]
		b = b[:l]

		var lock sync.Mutex

		theJob := &job{
			name: "Dot Product of float32",
			theFunc: func(pos int) bool {
				res := false
				r := int32(0)
				pos = pos * _MinWorkPerThread

				if pos < len(a) && pos < len(b) {
					aArr := a[pos:]
					bArr := b[pos:]
					l := gomath.MinInt(_MinWorkPerThread, gomath.MinInt(len(aArr), len(bArr)))
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

		parallels.GetRuntime().StartJob(theJob, false)
	}
	return result
}

// DotInt64 performs a dot product on two arrays of float64
func DotInt64(a, b []int64) int64 {
	result := int64(0)
	if a != nil && b != nil {
		l := gomath.MinInt(len(a), len(b))

		a = a[:l]
		b = b[:l]

		var lock sync.Mutex

		theJob := &job{
			name: "Dot Product of float32",
			theFunc: func(pos int) bool {
				res := false
				r := int64(0)
				pos = pos * _MinWorkPerThread

				if pos < len(a) && pos < len(b) {
					aArr := a[pos:]
					bArr := b[pos:]
					l := gomath.MinInt(_MinWorkPerThread, gomath.MinInt(len(aArr), len(bArr)))
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

		parallels.GetRuntime().StartJob(theJob, false)
	}
	return result
}
