package parallels

import "github.com/deathly809/gomath"

const (
	_BatchSize = 10000
)

// Foreach mimics a parallel foreach construct
//
//      This code should be equivalent to the serial code:
//
//          for i := 0 ; i < iterations; i++ {
//              f(i);
//          }
//
//  @f          -   Function to call in each iteration of the for loop
//  @iterations -   How many iterations we should perform
//
//
//
func Foreach(f func(i int) bool, iterations int) {
	rt := GetRuntime()
	rt.Start()

	jobs := []Job(nil)

	for iter := 0; iter < iterations; iter += _BatchSize {

		function := func(start, end int) func(int) bool {
			return func(i int) bool {
				if i >= end {
					return false
				}
				return f(i)
			}
		}(iter, gomath.MinInt(iter+_BatchSize, iterations))

		theJob := CreateIterableJob("Foreach", function, iter)

		jobs = append(jobs, theJob)
	}
	Run(jobs)
}
