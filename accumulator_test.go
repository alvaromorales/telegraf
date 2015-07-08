package telegraf

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func FakeMeasurement() (measurement string, fields map[string]interface{}, tags map[string]string) {
	measurement = "test_measurement"
	fields = map[string]interface{}{
		"value": "test",
	}
	tags = map[string]string{
		"tag1": "test",
	}
	return
}

func FakeMeasurementWithTime() (measurement string, fields map[string]interface{}, tags map[string]string, timestamp time.Time) {
	measurement, fields, tags = FakeMeasurement()
	timestamp = time.Now().AddDate(-1, 0, 0)
	return
}

func CheckAddedPoint(t *testing.T, bp *BatchPoints, measurement string, fields map[string]interface{}, tags map[string]string) {
	p := bp.Points[len(bp.Points)-1]
	assert.Equal(t, measurement, p.Measurement, "correct measurement (with prefix) should be passed")
	assert.Equal(t, fields, p.Fields, "correct value should be passed")
	assert.Equal(t, tags, p.Tags, "correct tags should be passed")
}

func CheckAddedPointWithTime(t *testing.T, bp *BatchPoints, measurement string, fields map[string]interface{}, tags map[string]string, timestamp time.Time) {
	CheckAddedPoint(t, bp, measurement, fields, tags)
	p := bp.Points[len(bp.Points)-1]
	assert.Equal(t, timestamp, p.Time, "correct timestamp should be passed")
}

func TestAdd(t *testing.T) {
	bp := new(BatchPoints)
	measurement, fields, tags := FakeMeasurement()
	bp.Add(measurement, fields["value"], tags)
	CheckAddedPoint(t, bp, measurement, fields, tags)
}

func TestAddWithTime(t *testing.T) {
	bp := new(BatchPoints)
	prefix := "testplugin"
	bp.Prefix = prefix
	measurement, fields, tags, timestamp := FakeMeasurementWithTime()
	bp.AddWithTime(measurement, fields["value"], tags, timestamp)
	CheckAddedPointWithTime(t, bp, prefix+measurement, fields, tags, timestamp)
}

func TestAddValues(t *testing.T) {
	bp := new(BatchPoints)
	prefix := "testplugin"
	bp.Prefix = prefix
	measurement, fields, tags := FakeMeasurement()
	fields["value2"] = 123
	bp.AddValues(measurement, fields, tags)
	CheckAddedPoint(t, bp, prefix+measurement, fields, tags)
}

func TestAddValuesWithTime(t *testing.T) {
	bp := new(BatchPoints)
	prefix := "testplugin"
	bp.Prefix = prefix
	measurement, fields, tags, timestamp := FakeMeasurementWithTime()
	fields["value2"] = 123
	bp.AddValuesWithTime(measurement, fields, tags, timestamp)
	CheckAddedPointWithTime(t, bp, prefix+measurement, fields, tags, timestamp)
}
