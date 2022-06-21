package fixtures

import m "github.com/RedHatInsights/sources-api-go/model"

var TestApplicationTypeData = []m.ApplicationType{
	{
		Id:          1,
		Name:        "app type name",
		DisplayName: "test app type",
	},
	{
		Id:          2,
		Name:        "second app type name",
		DisplayName: "second test app type",
	},
}
