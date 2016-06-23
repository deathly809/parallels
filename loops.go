package parallels

const (
	_BatchSize = 10000
)

// ForEach mimics a parallel foreach construct
func ForEach(f func(i int) bool, iterations int) {
	rt := GetRuntime()
	rt.Start()

	jobs := []Job(nil)

	for i := 0; i < iterations; i += _BatchSize {
		theJob := &job{
			id:      i,
			name:    "Foreach",
			pos:     i,
			theFunc: f,
		}
		jobs = append(jobs, theJob)
	}
	run(jobs)
}
