package generator

type ApiTest struct {
	Apis []Api
}

type Api struct {
	PackageName string
	FileName    string
	TestCases   []TestCase
}

type TestCase struct {
	Name      string
	Responses []string
}
