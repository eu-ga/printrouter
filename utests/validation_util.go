package utests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// BaseValidationTest basic structure of an request validation test
type BaseValidationTest struct {
	Name        string
	Req         ValidableRequest
	ExpectedErr error
}
type singleValidationTest struct {
	BaseValidationTest
	TestMock
}
type internalT struct {
	base BaseValidationTest
	t    *testing.T
}

// GetTestName return test name
func (inst internalT) GetTestName() string {
	return inst.base.Name
}

// ExecuteTest execute test and return an eventual error
func (inst internalT) ExecuteTest() error {
	return validationTest(inst.base, inst.t)
}

func validationTest(tc BaseValidationTest, t *testing.T) error {
	err := tc.Req.Validate()
	if tc.ExpectedErr == nil {
		assert.Equal(t, tc.ExpectedErr, err)
	} else {
		assert.Equal(t, tc.ExpectedErr.Error(), err.Error())
	}
	return err
}
func transformSingleBaseToValidationTest(base BaseValidationTest, t *testing.T) singleValidationTest {
	tMock := internalT{base, t}
	valTest := singleValidationTest{
		base, tMock,
	}
	return valTest
}

func transformAllBaseToValidationTest(base []BaseValidationTest, t *testing.T) []TestMock {
	validationTests := []TestMock{}
	for i := range base {
		validationTests = append(validationTests, transformSingleBaseToValidationTest(base[i], t))
	}
	return validationTests
}

// ExecValidationTest run validation tests between the expected request format and the mock input
func ExecValidationTest(base []BaseValidationTest, t *testing.T) []error {
	valt := transformAllBaseToValidationTest(base, t)
	return RunTest(valt, t)
}
