package histogram

import (
	"encoding/json"
	"log"

	"github.com/codahale/hdrhistogram"
)

// Histogram represents responce time histogram
type Histogram struct {
	histogram *hdrhistogram.Histogram
}

// New return `goad/histogram`
func New() *Histogram {
	maxNanoSeconds := int64(60 * 5 * 1000 * 1000)
	sigfigs := 3
	innerHist := hdrhistogram.New(0, maxNanoSeconds, sigfigs)

	return &Histogram{innerHist}
}

// RecordValue records the given value, return an error if the value is out of range
func (h *Histogram) RecordValue(v int64) error {
	return h.histogram.RecordValue(v)
}

// Export snapshot
func (h *Histogram) Export() *Snapshot {
	snap := h.histogram.Export()

	return &Snapshot{snap}
}

// Snapshot represents exported histogram as slice
type Snapshot struct {
	snapshot *hdrhistogram.Snapshot
}

func (s *Snapshot) String() string {
	js, err := json.Marshal(s)
	if err != nil {
		log.Fatal("Failed to marshal hdr.Snapshot")
	}

	return string(js)
}
