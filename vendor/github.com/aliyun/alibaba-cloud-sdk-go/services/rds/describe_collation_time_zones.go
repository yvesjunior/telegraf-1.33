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

// DescribeCollationTimeZones invokes the rds.DescribeCollationTimeZones API synchronously
func (client *Client) DescribeCollationTimeZones(request *DescribeCollationTimeZonesRequest) (response *DescribeCollationTimeZonesResponse, err error) {
	response = CreateDescribeCollationTimeZonesResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeCollationTimeZonesWithChan invokes the rds.DescribeCollationTimeZones API asynchronously
func (client *Client) DescribeCollationTimeZonesWithChan(request *DescribeCollationTimeZonesRequest) (<-chan *DescribeCollationTimeZonesResponse, <-chan error) {
	responseChan := make(chan *DescribeCollationTimeZonesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeCollationTimeZones(request)
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

// DescribeCollationTimeZonesWithCallback invokes the rds.DescribeCollationTimeZones API asynchronously
func (client *Client) DescribeCollationTimeZonesWithCallback(request *DescribeCollationTimeZonesRequest, callback func(response *DescribeCollationTimeZonesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeCollationTimeZonesResponse
		var err error
		defer close(result)
		response, err = client.DescribeCollationTimeZones(request)
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

// DescribeCollationTimeZonesRequest is the request struct for api DescribeCollationTimeZones
type DescribeCollationTimeZonesRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeCollationTimeZonesResponse is the response struct for api DescribeCollationTimeZones
type DescribeCollationTimeZonesResponse struct {
	*responses.BaseResponse
	RequestId          string             `json:"RequestId" xml:"RequestId"`
	CollationTimeZones CollationTimeZones `json:"CollationTimeZones" xml:"CollationTimeZones"`
}

// CreateDescribeCollationTimeZonesRequest creates a request to invoke DescribeCollationTimeZones API
func CreateDescribeCollationTimeZonesRequest() (request *DescribeCollationTimeZonesRequest) {
	request = &DescribeCollationTimeZonesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeCollationTimeZones", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeCollationTimeZonesResponse creates a response to parse from DescribeCollationTimeZones response
func CreateDescribeCollationTimeZonesResponse() (response *DescribeCollationTimeZonesResponse) {
	response = &DescribeCollationTimeZonesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}