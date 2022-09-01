package vms

type Categories struct {
	Name  string
	Value string
}

var Category1 = Categories{
	Name:  "2FA",
	Value: "2fa",
}

var Category2 = Categories{
	Name:  "New Infrastructure",
	Value: "new infrastructure",
}

var Category3 = Categories{
	Name:  "Permissions",
	Value: "permissions",
}

var Category4 = Categories{
	Name:  "Others",
	Value: "others",
}

var SupportedCategories = [...]Categories{
	Category1,
	Category2,
	Category3,
	Category4,
}
