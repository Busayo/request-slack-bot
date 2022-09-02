package generalrequests

type Urgency struct {
	Name  string
	Value string
}

var Urgency1 = Urgency{
	Name:  "High",
	Value: "high",
}

var Urgency2 = Urgency{
	Name:  "Medium",
	Value: "medium",
}

var Urgency3 = Urgency{
	Name:  "Low",
	Value: "low",
}



var SupportedUrgency= [...]Urgency{
	Urgency1,
	Urgency2,
	Urgency3,
}
