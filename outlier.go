package anomaly

import "time"

// Outlier represents dataset.
type Outlier struct {
	PeriodStart time.Time `json:"OutlierPeriodStart"`
	PeriodEnd   time.Time `json:"OutlierPeriodEnd"`
	Metric      string    `json:"Metric"`
	Attribute   string    `json:"Attribute"`

	// Mock dataset input data into Value property.
	// We will use it for detection based on 3 sigma method.
	Value float64 `json:"-"`
}
