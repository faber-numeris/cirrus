package tools

func GenerateARN(
	service string,
	region string,
	accountID string,
	resourceType string,
	resourceName string,
) string {
	return "arn:aws:" + service + ":" + region + ":" + accountID + ":" + resourceType + "/" + resourceName
}
