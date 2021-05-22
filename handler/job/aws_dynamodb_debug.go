// +build debug

package main

func registerToDynamoDB() error {
	log.Println(`[DEBUG] skip "registerToDynamoDB"`)

	return nil
}

func getAllFromDynamoDB() error {
	log.Println(`[DEBUG] skip "getAllFromDynamoDB"`)

	return nil
}

func deleteAllOfDynamoDB() error {
	log.Println(`[DEBUG] skip "deleteAllOfDynamoDB"`)

	return nil
}
