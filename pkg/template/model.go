package template

// GeneratedTest represents all the tests in Swagger Json file
type GeneratedTest struct {
	GeneratedTests map[string]Test // All the test cases in the swagger file
}

type Test struct {
	PackageName string     // Package Name for the generated test file
	FileName    string     // File Name for the generated test file
	TestCases   []TestCase // Test Cases in the single generated test file
}

type TestCase struct {
	Name         string // Name of the test case
	Description  string // Description of the test case
	Endpoint     string // Endpoint used for the test case
	Method       string // Method for the test case
	ResponseCode string // Response code expected for test case
}
