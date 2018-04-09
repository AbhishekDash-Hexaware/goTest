package e2etests

import (
	"testing"
)

func TestEmployeeDetails(t *testing.T) {
	var happyTests = []TestCase{
		TestCase{
			action:         "directname",
			utterances:     []string{"show details of krishna kumar"},
			parameters:     map[string]string{"employee_detail": "", "empname": "krishna kumar"},
			contains:       []string{"Krishna Kumar", "krishnakumar@hexaware.com", "36591", "CTO Office", "23-01-2017", "52505", "SIRUSERI 3", "9840984124"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "directid",
			utterances:     []string{"show details of 36591"},
			parameters:     map[string]string{"employee_detail": "", "empid": "36591"},
			contains:       []string{"Krishna Kumar", "krishnakumar@hexaware.com", "36591", "CTO Office", "23-01-2017", "52505", "SIRUSERI 3", "9840984124"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "usrinp_name",
			utterances:     []string{"I need an employee's details", "krishna kumar"},
			parameters:     map[string]string{"employee_detail": "", "empname": "krishna kumar"},
			contains:       []string{"Krishna Kumar", "krishnakumar@hexaware.com", "36591", "CTO Office", "23-01-2017", "52505", "SIRUSERI 3", "9840984124"},
			doesNotContain: []string{},
		},
		TestCase{
			action:         "usrinp_id",
			utterances:     []string{"I need an employee's details", "I think the id is 36591"},
			parameters:     map[string]string{"employee_detail": "", "empid": "36591"},
			contains:       []string{"Krishna Kumar", "krishnakumar@hexaware.com", "36591", "CTO Office", "23-01-2017", "52505", "SIRUSERI 3", "9840984124"},
			doesNotContain: []string{},
		},
	}
	runTestCases(t, happyTests)
}

func TestEmployeeSpecificDetail(t *testing.T) {
	var happyTests = []TestCase{
		TestCase{
			action:         "directname",
			utterances:     []string{"what is the cell number of krishna kumar?"},
			parameters:     map[string]string{"employee_detail": "mobile", "empname": "krishna kumar"},
			contains:       []string{"Krishna Kumar", "9840984124"},
			doesNotContain: []string{"36591", "CTO Office", "23-01-2017", "krishnakumar@hexaware.com", "SIRUSERI 3", "52505"},
		},
		TestCase{
			action:         "directid",
			utterances:     []string{"get the extension of 36591"},
			parameters:     map[string]string{"employee_detail": "extension", "empid": "36591"},
			contains:       []string{"Krishna Kumar", "52505"},
			doesNotContain: []string{"36591", "CTO Office", "23-01-2017", "krishnakumar@hexaware.com", "SIRUSERI 3", "9840984124"},
		},
		TestCase{
			action:         "directname",
			utterances:     []string{"where does krishna kumar sit?"},
			parameters:     map[string]string{"employee_detail": "location", "empname": "krishna kumar"},
			contains:       []string{"Krishna Kumar", "SIRUSERI 3"},
			doesNotContain: []string{"36591", "CTO Office", "23-01-2017", "krishnakumar@hexaware.com", "52505", "9840984124"},
		},
		TestCase{
			action:         "directid",
			utterances:     []string{"email address of 36591"},
			parameters:     map[string]string{"employee_detail": "email", "empid": "36591"},
			contains:       []string{"Krishna Kumar", "krishnakumar@hexaware.com"},
			doesNotContain: []string{"36591", "CTO Office", "23-01-2017", "52505", "SIRUSERI 3", "9840984124"},
		},
		TestCase{
			action:         "usrinp_id",
			utterances:     []string{"I need an employee's email address", "I think the id is 36591"},
			parameters:     map[string]string{"employee_detail": "email", "empid": "36591"},
			contains:       []string{"Krishna Kumar", "krishnakumar@hexaware.com"},
			doesNotContain: []string{"36591", "CTO Office", "23-01-2017", "52505", "SIRUSERI 3", "9840984124"},
		},
	}
	runTestCases(t, happyTests)
}
