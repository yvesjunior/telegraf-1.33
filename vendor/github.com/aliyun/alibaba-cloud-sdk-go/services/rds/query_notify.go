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

// QueryNotify invokes the rds.QueryNotify API synchronously
func (client *Client) QueryNotify(request *QueryNotifyRequest) (response *QueryNotifyResponse, err error) {
	response = CreateQueryNotifyResponse()
	err = client.DoAction(request, response)
	return
}

// QueryNotifyWithChan invokes the rds.QueryNotify API asynchronously
func (client *Client) QueryNotifyWithChan(request *QueryNotifyRequest) (<-chan *QueryNotifyResponse, <-chan error) {
	responseChan := make(chan *QueryNotifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryNotify(request)
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

// QueryNotifyWithCallback invokes the rds.QueryNotify API asynchronously
func (client *Client) QueryNotifyWithCallback(request *QueryNotifyRequest, callback func(response *QueryNotifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryNotifyResponse
		var err error
		defer close(result)
		response, err = client.QueryNotify(request)
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

// QueryNotifyRequest is the request struct for api QueryNotify
type QueryNotifyRequest struct {
	*requests.RpcRequest
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
	WithConfirmed requests.Boolean `position:"Body" name:"WithConfirmed"`
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	From          string           `position:"Body" name:"From"`
	To            string           `position:"Body" name:"To"`
}

// QueryNotifyResponse is the response struct for api QueryNotify
type QueryNotifyResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateQueryNotifyRequest creates a request to invoke QueryNotify API
func CreateQueryNotifyRequest() (request *QueryNotifyRequest) {
	request = &QueryNotifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "QueryNotify", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryNotifyResponse creates a response to parse from QueryNotify response
func CreateQueryNotifyResponse() (response *QueryNotifyResponse) {
	response = &QueryNotifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}