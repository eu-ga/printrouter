package utests

import (
	"testing"
)

// ValidableRequest interface that defines an validable request
type ValidableRequest interface {
	Validate() error
}

// TestMock interface that defines that an test can be run using the RunTest function
type TestMock interface {
	GetTestName() string
	ExecuteTest() error
}

// RunTest tests runner and error getter
func RunTest(tt []TestMock, t *testing.T) []error {
	errors := []error{}
	for i := range tt {
		t.Run(tt[i].GetTestName(), func(t *testing.T) {
			err := tt[i].ExecuteTest()
			errors = append(errors, err)
		})
	}
	return errors
}
