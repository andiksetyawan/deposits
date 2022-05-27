package threshold

import (
	"deposits/config"
	"deposits/model/view"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsAboveThreshold(t *testing.T) {
	now := time.Now()
	cs := &view.Threshold{
		WalletID:       "1",
		AboveThreshold: true,
		Histories: []view.History{
			{
				CreatedAt: now.Add(-time.Minute * 1),
				Amount:    6000,
			},
			{
				CreatedAt: now.Add(-time.Second * 5),
				Amount:    6000,
			},
		},
	}
	p := NewThresholdProcessor(config.GroupThreshold)
	aboveThreshold := p.isAboveThreshold(cs)
	assert.Equal(t, aboveThreshold, cs.AboveThreshold)
}

func TestThresholdFullCase(t *testing.T) {
	now := time.Now()
	cases := []struct {
		threshold view.Threshold
		expect    bool
	}{
		{
			expect: true,
			threshold: view.Threshold{
				WalletID:       "1",
				AboveThreshold: true,
				Histories: []view.History{
					{
						CreatedAt: now.Add(-time.Minute * 1),
						Amount:    6000,
					},
					{
						CreatedAt: now.Add(-time.Second * 5),
						Amount:    6000,
					},
				},
			},
		},
		{
			expect: false,
			threshold: view.Threshold{
				WalletID:       "1",
				AboveThreshold: false,
				Histories: []view.History{
					{
						CreatedAt: now.Add(-time.Minute * 3),
						Amount:    6000,
					},
					{
						CreatedAt: now.Add(-time.Second * 1),
						Amount:    6000,
					},
				},
			},
		},
		{
			expect: false,
			threshold: view.Threshold{
				WalletID:       "1",
				AboveThreshold: false,
				Histories: []view.History{
					{
						CreatedAt: now.Add(-time.Minute * 3),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 45),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 35),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 25),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 15),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 5),
						Amount:    2000,
					},
				},
			},
		},
		{
			expect: true,
			threshold: view.Threshold{
				WalletID:       "1",
				AboveThreshold: true,
				Histories: []view.History{
					{
						CreatedAt: now.Add(-time.Second * 100),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 45),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 35),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 25),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 15),
						Amount:    2000,
					},
					{
						CreatedAt: now.Add(-time.Second * 5),
						Amount:    2000,
					},
				},
			},
		},
	}

	for _, s := range cases {
		p := NewThresholdProcessor(config.GroupThreshold)
		aboveThreshold := p.isAboveThreshold(&s.threshold)
		assert.Equal(t, aboveThreshold, s.expect)
	}
}
