package vms

type Categories struct {
	Name  string
	Value string
}

var Category1 = Categories{
	Name:  "2FA",
	Value: "2FA",
}

var Category2 = Categories{
	Name:  "New Infrastructure",
	Value: "New Infrastructure",
}

var Category3 = Categories{
	Name:  "Permissions",
	Value: "Permissions",
}

var Category4 = Categories{
	Name:  "Others",
	Value: "Others",
}

var SupportedCategories = [...]Categories{
	Category1,
	Category2,
	Category3,
	Category4,
}
