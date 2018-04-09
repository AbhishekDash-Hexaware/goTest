package e2etests

import (
	"testing"
)

func TestBusDetails(t *testing.T) {
	var happyTests = []TestCase{
		TestCase{
			action:         "directbusno",
			utterances:     []string{"bus details of 609 in the morning"},
			parameters:     map[string]string{"busno": "609", "session": "morning"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"bus route of 609 in the evening"},
			parameters:     map[string]string{"busno": "609", "session": "evening"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"bus timings of 609 in the morning"},
			parameters:     map[string]string{"busno": "609", "session": "morning"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"I want the bus details of 609", "morning"},
			parameters:     map[string]string{"busno": "609", "session": "morning"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"please tell me the bus route of 609", "evening"},
			parameters:     map[string]string{"busno": "609", "session": "evening"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"what are the bus timings of 609", "morning"},
			parameters:     map[string]string{"busno": "609", "session": "morning"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"show me morning bus details", "609"},
			parameters:     map[string]string{"busno": "609", "session": "morning"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"give the evening bus route", "609"},
			parameters:     map[string]string{"busno": "609", "session": "evening"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusno",
			utterances:     []string{"tell me morning bus timings", "609"},
			parameters:     map[string]string{"busno": "609", "session": "morning"},
			contains:       []string{"CMBT", "MMDA", "Vadapalani", "K K Nagar", "Nesapakkam", "ESI Hospital", "Siruseri"},
			doesNotContain: []string{},
		}}
	runTestCases(t, happyTests)
}

func TestBusDetailsByStop(t *testing.T) {
	var happyTests = []TestCase{
		TestCase{
			action:         "directbusquestion",
			utterances:     []string{"what bus should I take to go to adyar"},
			parameters:     map[string]string{"location": "adyar", "session": "evening"},
			contains:       []string{"605"},
			doesNotContain: []string{},
		},
		// "more" works here since we are in the same session
		TestCase{
			action:         "directbusno",
			utterances:     []string{"more"},
			parameters:     map[string]string{"busno": "605", "session": "evening"},
			contains:       []string{"Chemmancheri Bus Stop", "AKDR Golf Village", "Mettukuppam MTC Bus Stop", "Thoraipakkam Indian Oil Petrolbunk", "Seevaram MTC Bus Stop", "Perungudi - Opp to Life line Hospital", "Kandhanchavadi MTC Bus Stop", "Madhya Kailash", "Anna Library (Kotturpuram)", "SIET", "GRT", "HT1", "Valluvarkottam", "Chetpet Signal", "KMC Bus Stand", "Ega Theatre"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusquestion",
			utterances:     []string{"which bus comes via velachery?"},
			parameters:     map[string]string{"location": "velachery", "session": "morning"},
			contains:       []string{"602"},
			doesNotContain: []string{},
		},
		// "more" works here since we are in the same session
		TestCase{
			action:         "directbusno",
			utterances:     []string{"more"},
			parameters:     map[string]string{"busno": "602", "session": "morning"},
			contains:       []string{"Ragavendra Kalyana Mandapam", "Five Lights", "Ayodhya Mandapam", "West Mambalam SRM Hospital", "Postal Colony", "Srinivasa Theatre", "Aranganathar Subway", "Saidapet Court Bus Stop", "Velachery Check Post", "Velachery Lake Bus Stop", "Velachery Murugan Mandabam", "Velachery Railway Station", "Siruseri"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusquestion",
			utterances:     []string{"what bus should I take?", "coming back from the office", "adyar"},
			parameters:     map[string]string{"location": "adyar", "session": "evening"},
			contains:       []string{"605"},
			doesNotContain: []string{},
		},
		// "more" works here since we are in the same session
		TestCase{
			action:         "directbusno",
			utterances:     []string{"more"},
			parameters:     map[string]string{"busno": "605", "session": "evening"},
			contains:       []string{"Chemmancheri Bus Stop", "AKDR Golf Village", "Mettukuppam MTC Bus Stop", "Thoraipakkam Indian Oil Petrolbunk", "Seevaram MTC Bus Stop", "Perungudi - Opp to Life line Hospital", "Kandhanchavadi MTC Bus Stop", "Madhya Kailash", "Anna Library (Kotturpuram)", "SIET", "GRT", "HT1", "Valluvarkottam", "Chetpet Signal", "KMC Bus Stand", "Ega Theatre"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directbusquestion",
			utterances:     []string{"I want to take a bus", "going to the office", "velachery"},
			parameters:     map[string]string{"location": "velachery", "session": "morning"},
			contains:       []string{"602"},
			doesNotContain: []string{},
		},
		// "more" works here since we are in the same session
		TestCase{
			action:         "directbusno",
			utterances:     []string{"more"},
			parameters:     map[string]string{"busno": "602", "session": "morning"},
			contains:       []string{"Ragavendra Kalyana Mandapam", "Five Lights", "Ayodhya Mandapam", "West Mambalam SRM Hospital", "Postal Colony", "Srinivasa Theatre", "Aranganathar Subway", "Saidapet Court Bus Stop", "Velachery Check Post", "Velachery Lake Bus Stop", "Velachery Murugan Mandabam", "Velachery Railway Station", "Siruseri"},
			doesNotContain: []string{},
		},
	}
	runTestCases(t, happyTests)
}
