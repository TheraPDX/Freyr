package fake

import (
	"errors"
	"github.com/serdmanczyk/freyr/models"
	"math/rand"
	"time"
)

var randGen *rand.Rand

func init() {
	randGen = rand.New(rand.NewSource(time.Now().Unix()))
}

func floatBetween(a, b float32) float32 {
	if b <= a {
		return 0.0
	}
	return (randGen.Float32() * (b - a)) + a
}

func RandReading(userEmail, core string, posted time.Time) models.Reading {
	return models.Reading{
		UserEmail:   userEmail,
		CoreId:      core,
		Posted:      posted,
		Temperature: floatBetween(10.0, 20.0),
		Humidity:    floatBetween(30.0, 60.0),
		Moisture:    floatBetween(20.0, 96.0),
		Light:       floatBetween(0.0, 120.0),
		Battery:     floatBetween(0.0, 100.0),
	}
}

type ReadingStore struct {
	readings []models.Reading
}

func (f *ReadingStore) StoreReading(reading models.Reading) error {
	f.readings = append(f.readings, reading)
	return nil
}

func (f *ReadingStore) GetLatestReadings(userEmail string) (readings []models.Reading, err error) {
	err = errors.New("don't need this yet")
	return
}

func (f *ReadingStore) GetReadings(core string, start, end time.Time) ([]models.Reading, error) {
	filtered := models.FilterReadings(f.readings, func(r models.Reading) bool {
		if r.CoreId == core && r.Posted.After(start) && r.Posted.Before(end) {
			return true
		}

		return false
	})

	return filtered, nil
}
