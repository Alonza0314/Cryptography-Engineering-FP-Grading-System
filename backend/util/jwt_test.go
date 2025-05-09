package util_test

import (
	"ce/backend/util"
	"testing"
	"time"
)

var testCasesJwtToken = []struct {
	studentId        string
	ioa              int64
	exp              int64
	alreadyTotp      bool
	expectedResult   bool
	expectedError    error
}{
	{
		studentId:        "1234567890",
		ioa:              time.Now().Unix(),
		exp:              time.Now().Unix() + 1000,
		alreadyTotp:      true,
		expectedResult:   true,
		expectedError:    nil,
	},
}

func TestJwtToken(t *testing.T) {
	for _, testCase := range testCasesJwtToken {
		jwtToken, err := util.GenerateJwtToken(testCase.studentId, testCase.ioa, testCase.exp, testCase.alreadyTotp)
		if err != nil {
			t.Errorf("Error generating JWT token: %v", err)
		}
		
		result, studentId, err := util.VerifyJwtToken(jwtToken)
		if err != nil {
			t.Errorf("Error verifying JWT token: %v", err)
		}
		if result != testCase.expectedResult {
			t.Errorf("Expected result: %t, but got: %t", testCase.expectedResult, result)
		}
		if studentId != testCase.studentId {
			t.Errorf("Expected studentId: %s, but got: %s", testCase.studentId, studentId)
		}
	}
}
