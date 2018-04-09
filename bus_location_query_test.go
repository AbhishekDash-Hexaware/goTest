package e2etests

import (
	"testing"
)

func TestBusLocation(t *testing.T) {
	var happyTests = []TestCase{
		TestCase{
			action:         "buslocation",
			utterances:     []string{"where is 609?"},
			parameters:     map[string]string{"busno": "609"},
			contains:       []string{"Chennai"},
			doesNotContain: []string{"Oops"},
		},
		TestCase{
			action:         "buslocation",
			utterances:     []string{"where is 609 now"},
			parameters:     map[string]string{"busno": "609"},
			contains:       []string{"Chennai"},
			doesNotContain: []string{"Oops"},
		},
		TestCase{
			action:         "buslocation",
			utterances:     []string{"give me the location of 609"},
			parameters:     map[string]string{"busno": "609"},
			contains:       []string{"Chennai"},
			doesNotContain: []string{"Oops"},
		},
		TestCase{
			action:         "buslocation",
			utterances:     []string{"locate 609 now"},
			parameters:     map[string]string{"busno": "609"},
			contains:       []string{"Chennai"},
			doesNotContain: []string{"Oops"},
		},
		TestCase{
			action:         "buslocation",
			utterances:     []string{"locate bus", "609"},
			parameters:     map[string]string{"busno": "609"},
			contains:       []string{"Chennai"},
			doesNotContain: []string{"Oops"},
		},
		TestCase{
			action:         "buslocation",
			utterances:     []string{"locate 606 now"},
			parameters:     map[string]string{"busno": "606"},
			contains:       []string{"Oops"},
			doesNotContain: []string{"Chennai"},
		},
		TestCase{
			action:         "buslocation",
			utterances:     []string{"locate bus", "606"},
			parameters:     map[string]string{"busno": "606"},
			contains:       []string{"Oops"},
			doesNotContain: []string{"Chennai"},
		},
	}
	runTestCases(t, happyTests)
}
