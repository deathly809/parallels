package parallels

const (
	_BatchSize = 10000
)

// Foreach mimics a parallel foreach construct
func Foreach(f func(i int) bool, iterations int) {
	rt := GetRuntime()
	rt.Start()

	jobs := []Job(nil)

	for i := 0; i < iterations; i += _BatchSize {
		theJob := &IterableJob{
			ID:      i,
			Name:    "Foreach",
			Pos:     i,
			TheFunc: f,
		}
		jobs = append(jobs, theJob)
	}
	Run(jobs)
}
