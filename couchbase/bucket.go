package couchbase

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/couchbase/gocb/v2"
)

// getBucketConflictResolutionType custom function for get bucket conflict resolution type because couchbase golang sdk doesn't support to get conflict
// resolution type in gocb v2 version
func (cc *CouchbaseConnection) getBucketConflictResolutionType(bucketName string) (*gocb.ConflictResolutionType, error) {
	var conflictResolutionType conflictResolutionType
	var host string
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: cc.ClusterOptions.SecurityConfig.TLSSkipVerify},
	}

	client := http.Client{
		Timeout:   cc.ClusterOptions.TimeoutsConfig.ManagementTimeout,
		Transport: tr,
	}

	// TODO
	s := strings.Split(cc.ConnStr, "://")
	if s[0] == "couchbases" {
		host = fmt.Sprintf("https://%s:18091", s[1])
	} else {
		host = fmt.Sprintf("http://%s:8091", s[1])
	}

	// TODO https
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/pools/default/buckets/%s", host, bucketName), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(cc.ClusterOptions.Username, cc.ClusterOptions.Password)

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(resData, &conflictResolutionType); err != nil {
		return nil, err
	}

	return &conflictResolutionType.ConflictResolutionType, nil
}
