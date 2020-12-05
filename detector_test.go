package anomaly

import (
	"testing"
	"time"
)

func TestSigmaDetector(t *testing.T) {
	testCases := []struct {
		Name             string
		History          []float64
		Options          *SigmaOptions
		Input            []Outlier
		ExpectedWarnings int
		ExpectedAlarms   int
	}{
		{
			Name:    "Case1",
			History: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Options: nil, // default one
			Input: []Outlier{
				{
					PeriodStart: time.Now().UTC().Add(-1 * time.Hour),
					PeriodEnd:   time.Now().UTC().Add(1 * time.Hour),
					Metric:      "Visits(qGaHiHepUD1pLl_cWoD9aQ)",
					Attribute:   "Device Type > Tablet > 10(eAeciHepUD1pLl_cWoD932)",
					Value:       2.9,
				},
				{
					PeriodStart: time.Now().UTC().Add(-1 * time.Hour),
					PeriodEnd:   time.Now().UTC().Add(1 * time.Hour),
					Metric:      "Visits(qGaHiHepUD1pLl_cWoD9aQ)",
					Attribute:   "Device Type > Tablet > 10(eAeciHepUD1pLl_cWoD932)",
					Value:       5.8,
				},
			},
			ExpectedWarnings: 1,
			ExpectedAlarms:   1,
		},
		{
			Name:    "Case2",
			History: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
			Options: &SigmaOptions{
				Multiplier:       2.5,
				StrongMultiplier: 4.0,
			},
			Input: []Outlier{
				{
					PeriodStart: time.Now().UTC().Add(-1 * time.Hour),
					PeriodEnd:   time.Now().UTC().Add(1 * time.Hour),
					Metric:      "Visits(qGaHiHepUD1pLl_cWoD9aQ)",
					Attribute:   "Device Type > Tablet > 10(eAeciHepUD1pLl_cWoD932)",
					Value:       2.9,
				},
				{
					PeriodStart: time.Now().UTC().Add(-1 * time.Hour),
					PeriodEnd:   time.Now().UTC().Add(1 * time.Hour),
					Metric:      "Visits(qGaHiHepUD1pLl_cWoD9aQ)",
					Attribute:   "Device Type > Tablet > 10(eAeciHepUD1pLl_cWoD932)",
					Value:       5.8,
				},
			},
			ExpectedWarnings: 2,
			ExpectedAlarms:   0,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			detector := NewSigma(tc.History, tc.Options)

			res, err := detector.Detect(tc.Input...)
			if err != nil {
				t.Fatalf("got error: %v", err)
			}

			if len(res.Warnings) != tc.ExpectedWarnings {
				t.Errorf("warnings: have %d want %d", len(res.Warnings), tc.ExpectedWarnings)
			}

			if len(res.Alarms) != tc.ExpectedAlarms {
				t.Errorf("alarms: have %d want %d", len(res.Alarms), tc.ExpectedAlarms)
			}
		})
	}
}
