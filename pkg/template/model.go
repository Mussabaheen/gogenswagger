package template

type GeneratedTest struct {
	GeneratedTests map[string]Test
}

type Test struct {
	PackageName string
	FileName    string
	TestCases   []TestCase
}

type TestCase struct {
	Name        string
	Description string
}
