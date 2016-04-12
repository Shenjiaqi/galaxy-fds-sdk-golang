package Model

import "github.com/bitly/go-simplejson"

type FDSObjectSummary struct {
	BucketName string
	ObjectName string
	Owner      Owner
	Size       int64
}

func NewFDSObjectSummary(jsonValue simplejson.Json) (*FDSObjectSummary, error) {
	bucketName, err := jsonValue.Get("bucketName").String()
	if err != nil {
		return nil, err
	}
	objectName, err := jsonValue.Get("objectName").String()
	if err != nil {
		return nil, err
	}
	owner, err      := NewOwner(jsonValue.Get("owner"))
	if err != nil {
		return nil, err
	}
	size, err       := jsonValue.Get("size").Int64()
	if err != nil {
		return nil, err
	}
	return &FDSObjectSummary{
		BucketName: bucketName,
		ObjectName: objectName,
		Owner:      owner,
		Size:       size,
	}, nil
}