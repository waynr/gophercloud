/*
Package quotasets enables retrieving and managing Block Storage quotas.

Example to Get a Quota Set

	quotaset, err := quotasets.Get(blockStorageClient, "project-id").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)

Example to Get a Detailed Quota Set

	quotaset, err := quotasets.GetDetail(blockStorageClient, "project-id").Extract()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", quotaset)
*/
package quotasets
