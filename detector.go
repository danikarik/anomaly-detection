package anomaly

// Result holds all outliers.
//
// Between mean and 2 sigma - Warning.
// Between 2 sigma and 3 sigma - Alarm.
type Result struct {
	Warnings []Outlier
	Alarms   []Outlier
}

// Detector represents `OutliersDetectionMethod`.
type Detector interface {
	Detect(in ...Outlier) (*Result, error)
}

type threeSigma struct {
	sd   float64
	sig2 float64
	sig3 float64
}

// SigmaOptions is used to configure 3 sigma based detector.
type SigmaOptions struct {
	Multiplier       float64 `json:"OutliersMultiplier"`
	StrongMultiplier float64 `json:"StrongOutliersMultiplier"`
}

// NewSigma returns a new detector based on 3 sigma method.
//
// For demo puspose we will use mock history.
func NewSigma(history []float64, opts *SigmaOptions) Detector {
	if opts == nil {
		// Set default values.
		opts = &SigmaOptions{
			Multiplier:       2.0,
			StrongMultiplier: 3.0,
		}
	}

	sd := calcSD(history)

	return &threeSigma{
		sd:   sd,
		sig2: opts.Multiplier * sd,
		sig3: opts.StrongMultiplier * sd,
	}
}

func (d *threeSigma) Detect(in ...Outlier) (*Result, error) {
	res := &Result{
		Warnings: make([]Outlier, 0),
		Alarms:   make([]Outlier, 0),
	}

	for _, outlier := range in {
		if d.isWarning(outlier.Value) {
			res.Warnings = append(res.Warnings, outlier)
			continue
		}

		if d.isAlarm(outlier.Value) {
			res.Alarms = append(res.Alarms, outlier)
		}
	}

	return res, nil
}

func (d *threeSigma) isWarning(in float64) bool {
	return in > d.sd && in <= d.sig2
}

func (d *threeSigma) isAlarm(in float64) bool {
	return in > d.sig2 && in <= d.sig3
}
