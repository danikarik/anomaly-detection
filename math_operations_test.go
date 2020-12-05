package anomaly

import "testing"

func TestCalcMean(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    []float64
		Expected float64
	}{
		{
			Name:     "Case1",
			Input:    []float64{1, 2, 3},
			Expected: 2,
		},
		{
			Name:     "Case2",
			Input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Expected: 5.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			mean := calcMean(tc.Input)
			if mean != tc.Expected {
				t.Errorf("have %f want %f", mean, tc.Expected)
			}
		})
	}
}

func TestCalcSD(t *testing.T) {
	testCases := []struct {
		Name     string
		Input    []float64
		Expected float64
	}{
		{
			Name:     "Case1",
			Input:    []float64{1, 2, 3},
			Expected: 0.816496580927726,
		},
		{
			Name:     "Case2",
			Input:    []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Expected: 2.8722813232690143,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			sd := calcSD(tc.Input)
			if sd != tc.Expected {
				t.Errorf("have %f want %f", sd, tc.Expected)
			}
		})
	}
}
