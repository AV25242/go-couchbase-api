package utils

import (
	"models"
	"sync"

	gocb "gopkg.in/couchbase/gocb.v1"
)

var globalCluster *gocb.Cluster = nil
var globalBucket *gocb.Bucket = nil

var instance models.Configuration
var once sync.Once

func Bucket() gocb.Bucket {
	once.Do(func() {
		var err error

		instance = models.GetConfiguration()
		globalCluster, err = gocb.Connect(instance.Connection)
		if err != nil {
			panic(err)
		}
		globalCluster.Authenticate(gocb.PasswordAuthenticator{
			instance.Username,
			instance.Password,
		})
		globalBucket, err = globalCluster.OpenBucket(instance.Bucket, "")
		if err != nil {
			panic(err)
		}

	})
	return *globalBucket
}
