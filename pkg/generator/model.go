package generator

type ApiTest struct {
	Apis map[string]Api
}

type Api struct {
	PackageName string
	FileName    string
	TestCases   []TestCase
}

type TestCase struct {
	Name        string
	Description string
}
