package rds

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeCloudMigrationResult invokes the rds.DescribeCloudMigrationResult API synchronously
func (client *Client) DescribeCloudMigrationResult(request *DescribeCloudMigrationResultRequest) (response *DescribeCloudMigrationResultResponse, err error) {
	response = CreateDescribeCloudMigrationResultResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCloudMigrationResultWithChan invokes the rds.DescribeCloudMigrationResult API asynchronously
func (client *Client) DescribeCloudMigrationResultWithChan(request *DescribeCloudMigrationResultRequest) (<-chan *DescribeCloudMigrationResultResponse, <-chan error) {
	responseChan := make(chan *DescribeCloudMigrationResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCloudMigrationResult(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeCloudMigrationResultWithCallback invokes the rds.DescribeCloudMigrationResult API asynchronously
func (client *Client) DescribeCloudMigrationResultWithCallback(request *DescribeCloudMigrationResultRequest, callback func(response *DescribeCloudMigrationResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCloudMigrationResultResponse
		var err error
		defer close(result)
		response, err = client.DescribeCloudMigrationResult(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeCloudMigrationResultRequest is the request struct for api DescribeCloudMigrationResult
type DescribeCloudMigrationResultRequest struct {
	*requests.RpcRequest
	DBInstanceName       string           `position:"Query" name:"DBInstanceName"`
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	TaskName             string           `position:"Query" name:"TaskName"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	SourcePort           requests.Integer `position:"Query" name:"SourcePort"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	TaskId               requests.Integer `position:"Query" name:"TaskId"`
	SourceIpAddress      string           `position:"Query" name:"SourceIpAddress"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCloudMigrationResultResponse is the response struct for api DescribeCloudMigrationResult
type DescribeCloudMigrationResultResponse struct {
	*responses.BaseResponse
	TotalSize  int     `json:"TotalSize" xml:"TotalSize"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	PageNumber int64   `json:"PageNumber" xml:"PageNumber"`
	PageSize   int64   `json:"PageSize" xml:"PageSize"`
	Items      []Tasks `json:"Items" xml:"Items"`
}

// CreateDescribeCloudMigrationResultRequest creates a request to invoke DescribeCloudMigrationResult API
func CreateDescribeCloudMigrationResultRequest() (request *DescribeCloudMigrationResultRequest) {
	request = &DescribeCloudMigrationResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeCloudMigrationResult", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCloudMigrationResultResponse creates a response to parse from DescribeCloudMigrationResult response
func CreateDescribeCloudMigrationResultResponse() (response *DescribeCloudMigrationResultResponse) {
	response = &DescribeCloudMigrationResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}