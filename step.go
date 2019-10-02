package allure

type Step struct {
	Name          string
	Status        string
	StatusDetails StatusDetails
	Stage         string
	Start         int64
	Stop          int64
	Attachements  []Attachment
}

// func Step(t *testing.T, name string, stepFunc func()) {

// }
