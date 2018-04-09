package e2etests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"
)

type Query struct {
	Query    string `json:"query,omitempty"`
	Contexts []struct {
		Name       string            `json:"name,omitempty"`
		Parameters map[string]string `json:"parameters,omitempty"`
		Lifespan   int               `json:"lifespan,omitempty"`
	} `json:"contexts,omitempty"`
	ResetContexts bool   `json:"resetContexts,omitempty"`
	SessionID     string `json:"sessionId,omitempty"`
	Lang          string `json:"lang,omitempty"`
}

func NewQuery() Query {
	q := new(Query)
	q.Query = ""
	q.Lang = "en"
	rand.Seed(time.Now().UnixNano())
	q.SessionID = strconv.Itoa(rand.Int())
	return *q
}

type Response struct {
	Status struct {
		Code         int    `json:"code"`
		ErrorType    string `json:"errorType"`
		ErrorID      string `json:"errorId"`
		ErrorDetails string `json:"errorDetails"`
	}
	Result struct {
		Source           string            `json:"Source"`
		Action           string            `json:"action"`
		ActionIncomplete bool              `json:"actionIncomplete"`
		Parameters       map[string]string `json:"Parameters"`
		Contexts         []struct {
			Name       string            `json:"name,omitempty"`
			Parameters map[string]string `json:"parameters,omitempty"`
			Lifespan   int               `json:"lifespan,omitempty"`
		} `json:"contexts"`
		Metadata struct {
			IntentID                  string `json:"intentId"`
			WebhookUsed               string `json:"webhookUsed"`
			WebhookForSlotFillingUsed string `json:"webhookForSlotFillingUsed"`
			IntentName                string `json:"intentName"`
		} `json:"metadata"`
		Fulfillment struct {
			Speech      string `json:"speech"`
			DisplayText string `json:"displayText"`
			Messages    []struct {
				Type int `json:"type"`
			} `json:"messages"`
		} `json:"fulfillment"`
	} `json:"result"`
}

func postQuery(query Query, t *testing.T) Response {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(query)
	//fmt.Println("Printing query", b)
	req, err := http.NewRequest("POST", "https://api.api.ai/v1/query?v=20150910", b)
	if err != nil {
		log.Println(err)
		t.Error("Post to api.ai failed with", err)
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("CLIENT_TOKEN"))
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Error("Post to api.ai failed with error: ", err)
	}
	defer resp.Body.Close()
	dec := json.NewDecoder(resp.Body)
	var response Response
	dec.Decode(&response)
	if resp.StatusCode != 200 {
		t.Error("Post to api.ai failed with statusCode: ", resp.StatusCode, ", status: ", resp.Status, ", details: ", response.Status.ErrorDetails)
	}
	return response
}

func assertEqualMessage(t *testing.T, expected interface{}, actual interface{}, message string) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	t.Fatal(message)
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if reflect.DeepEqual(expected, actual) {
		return
	}
	message := fmt.Sprintf("(expected) %v != (actual)  %v", expected, actual)
	t.Fatal(message)
}

func assertContains(t *testing.T, expected string, actual string) {
	if strings.Contains(actual, expected) {
		return
	}
	message := fmt.Sprintf("(expected) %v not in (actual)  %v", expected, actual)
	t.Fatal(message)
}

func assertNotContains(t *testing.T, expected string, actual string) {
	if !strings.Contains(actual, expected) {
		return
	}
	message := fmt.Sprintf("(expected) %v in (actual)  %v", expected, actual)
	t.Fatal(message)
}

//TestCase models simple test cases which can be executed in a data-driven manner by runTestCase
type TestCase struct {
	action         string
	utterances     []string
	parameters     map[string]string
	contains       []string
	doesNotContain []string
}

func runTestCases(t *testing.T, testCases []TestCase) {
	q := NewQuery()
	for _, test := range testCases {
		t.Run(test.utterances[0], func(t *testing.T) {
			for index, utterance := range test.utterances {
				q.Query = utterance
				r := postQuery(q, t)
				//j, _ := json.MarshalIndent(r, "", " ")
				//fmt.Println("DEBUG: response: ", string(j))

				if index < len(test.utterances)-1 {
					if r.Result.Action != "" && r.Result.ActionIncomplete == false {
						t.Fatal("Action complete even before all utterances are uttered")
					}
				} else {

					assertEqual(t, test.action, r.Result.Action)
					assertEqual(t, false, r.Result.ActionIncomplete)
					assertEqual(t, test.parameters, r.Result.Parameters)
					for _, shouldContain := range test.contains {
						assertContains(t, shouldContain, r.Result.Fulfillment.DisplayText)
					}
					for _, shouldNotContain := range test.doesNotContain {
						assertNotContains(t, shouldNotContain, r.Result.Fulfillment.DisplayText)
					}
				}
			}
		})
	}
}
