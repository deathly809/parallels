package vectors

import (
	"log"
	"math/rand"
	"os"
	"runtime/pprof"
	"testing"

	"github.com/deathly809/gomath"
)

const (
	NRR = 400000000
)

var (
	ARR = []float32(nil)
	BRR = []float32(nil)
)

func initBenchmark() {
	ARR = genFloat32Array(NRR, 0.0, 1.0)
	BRR = genFloat32Array(NRR, 0.0, 1.0)
}

const (
	LowFloat32  = float32(1E-4)
	HighFloat32 = float32(1E6)

	LowInt32  = -100000
	HighInt32 = 100000
)

func genFloat32Array(n int, min, max float32) []float32 {
	A := make([]float32, n)
	for i := 0; i < n; i++ {
		A[i] = gomath.ScaleFloat32(min, max, 0, 1.0, rand.Float32())
	}
	return A
}

func genInt32Array(n int, min, max float32) []float32 {
	A := make([]float32, n)
	for i := 0; i < n; i++ {
		A[i] = gomath.ScaleFloat32(min, max, 0, 1.0, rand.Float32())
	}
	return A
}

func TestDotFloat32WithSameLength(t *testing.T) {
	N := 500000

	a := genFloat32Array(N, LowFloat32, HighFloat32)
	b := genFloat32Array(N, LowFloat32, HighFloat32)

	Expected := 0.0
	for i := range a {
		Expected += float64(a[i] * b[i])
	}

	computed := DotFloat32(a, b)

	if gomath.AbsFloat32(computed-float32(Expected)) < LowFloat32 {
		t.Logf("Expected %f but computed %f\n", Expected, computed)
		t.FailNow()
	}

}

func TestDotFloat32WithDiffLength(t *testing.T) {
	N := 1000 + rand.Intn(1000000)
	M := 1000 + rand.Intn(1000000)
	if N == M {
		N++
	}
	a := genFloat32Array(N, LowFloat32, HighFloat32)
	b := genFloat32Array(M, LowFloat32, HighFloat32)
	Expected := 0.0
	for i := range a {
		if i < N && i < M {
			Expected += float64(a[i] * b[i])
		} else {
			break
		}
	}

	computed := DotFloat32(a, b)

	if gomath.AbsFloat32(computed-float32(Expected)) < LowFloat32 {
		t.Logf("Expected %f but computed %f\n", Expected, computed)
		t.FailNow()
	}
}

func TestDotInt32WithSameLength(t *testing.T) {
	N := 1000 + rand.Intn(1000000)
	a := make([]int32, N)
	b := make([]int32, N)
	Expected := int32(0)
	for i := range a {
		a[i] = rand.Int31n(1000)
		b[i] = rand.Int31n(1000)

		Expected += int32(a[i] * b[i])
	}

	computed := DotInt32(a, b)

	if computed != int32(Expected) {
		t.Logf("Expected %d but computed %d\n", Expected, computed)
		t.FailNow()
	}
}

func TestDotInt32WithDiffLength(t *testing.T) {
	N := 1000 + rand.Intn(1000000)
	M := 1000 + rand.Intn(1000000)
	a := make([]int32, N)
	b := make([]int32, M)
	expected64 := int64(0)
	for i := range a {
		if i < N {
			a[i] = gomath.ScaleInt32(LowInt32, HighInt32, 0, HighInt32, rand.Int31n(HighInt32))
		}
		if i < M {
			b[i] = gomath.ScaleInt32(LowInt32, HighInt32, 0, HighInt32, rand.Int31n(HighInt32))
		}

		if i < N && i < M {
			expected64 += int64(a[i] * b[i])
		}
	}

	Expected := int32(expected64)
	Computed := DotInt32(a, b)

	if Computed != Expected {
		t.Logf("Expected %d but computed %d\n", Expected, Computed)
		t.FailNow()
	}
}

func performOp(b *testing.B) {
	if ARR == nil {
		initBenchmark()
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := DotFloat32(ARR[:], BRR[:])
		if result == 0.0 {
			b.Log("Lolwut")
		}
	}

}

func BenchmarkDotFloat32WorkPerThread1000(b *testing.B) {
	_MinWorkPerThread = 1000
	performOp(b)
}

func BenchmarkDotFloat32WorkPerThread10000(b *testing.B) {
	_MinWorkPerThread = 10000
	performOp(b)
}

func BenchmarkDotFloat32WorkPerThread100000(b *testing.B) {
	_MinWorkPerThread = 100000
	performOp(b)
}

func BenchmarkDotFloat32WorkPerThread1000000(b *testing.B) {
	_MinWorkPerThread = 1000000
	performOp(b)
}

func BenchmarkDotFloat32WorkPerThread10000000(b *testing.B) {
	_MinWorkPerThread = 10000000
	performOp(b)
}

func BenchmarkDotFloat32WorkPerThread100000000(b *testing.B) {
	f, err := os.Create("profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()
	_MinWorkPerThread = 100000000
	performOp(b)
}

func BenchmarkDotCPU(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := float32(0)
		for j := 0; j < NRR; j++ {
			result += ARR[j] * BRR[j]
		}
		if result == 0.0 {
			b.Log("Lolwut")
		}
	}
}
