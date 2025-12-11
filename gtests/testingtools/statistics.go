// Andre R. R. Costa * github.com/andrerrcosta2 * andrerrcosta@gmail.com

package testingtools

import (
	"math"

	"github.com/andrerrcosta2/gtools/core/domain/constraints/prim/nums"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/config/testlogs"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/gtests"
	"github.com/andrerrcosta2/gtools/gtests/testingtools/internal/logs"
)

// StatsLite creates a testing tool instance for simple statistics tests.
func StatsLite[C, N nums.Real](
	testing gtests.FailureLoggableTesting, buckets int, min, max N) gtests.StatsLite[C, N] {
	return &statsLite[C, N]{
		levelLoggableLite: levelLoggableLite[gtests.FailureLoggableTesting]{
			FailureLoggableTesting: testing,
			loggerLevel:            testlogs.Default,
			logger:                 logs.TimerStack(),
			errors:                 0,
		},
		classes: make(map[C]N),
		size:    buckets,
		minimum: min,
		maximum: max,
	}
}

type statsLite[C, N nums.Real] struct {
	levelLoggableLite[gtests.FailureLoggableTesting]
	classes map[C]N
	size    int
	count   int
	minimum N
	maximum N
}

func (s *statsLite[C, N]) Add(value C, weight N) {
	s.count++
	s.classes[value] = s.classes[value] + weight
}

func (s *statsLite[C, N]) Entropy() float64 {
	total := float64(s.count)
	var entropy float64
	for _, weight := range s.classes {
		p := float64(weight) / total
		if p > 0 {
			entropy -= p * math.Log2(p)
		}
	}
	return entropy
}

func (s *statsLite[C, N]) IsUnderDeviationOf(d float64) bool {
	expected := float64(s.count) / float64(s.size)
	upper := expected * (1 + d)
	lower := expected * (1 - d)

	for _, weight := range s.classes {
		f := float64(weight)
		if f < lower || f > upper {
			return false
		}
	}
	return true
}

func (s *statsLite[C, N]) ChiSquare() float64 {
	expected := float64(s.count) / float64(s.size)
	var chiSq float64
	for _, weight := range s.classes {
		diff := float64(weight) - expected
		chiSq += (diff * diff) / expected
	}
	return chiSq
}

func (s *statsLite[C, N]) Reset() {
	s.count = 0
	s.classes = make(map[C]N)
}
