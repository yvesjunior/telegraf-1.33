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

// ReceiveDBInstance invokes the rds.ReceiveDBInstance API synchronously
func (client *Client) ReceiveDBInstance(request *ReceiveDBInstanceRequest) (response *ReceiveDBInstanceResponse, err error) {
	response = CreateReceiveDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// ReceiveDBInstanceWithChan invokes the rds.ReceiveDBInstance API asynchronously
func (client *Client) ReceiveDBInstanceWithChan(request *ReceiveDBInstanceRequest) (<-chan *ReceiveDBInstanceResponse, <-chan error) {
	responseChan := make(chan *ReceiveDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ReceiveDBInstance(request)
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

// ReceiveDBInstanceWithCallback invokes the rds.ReceiveDBInstance API asynchronously
func (client *Client) ReceiveDBInstanceWithCallback(request *ReceiveDBInstanceRequest, callback func(response *ReceiveDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ReceiveDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.ReceiveDBInstance(request)
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

// ReceiveDBInstanceRequest is the request struct for api ReceiveDBInstance
type ReceiveDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	GuardDBInstanceId    string           `position:"Query" name:"GuardDBInstanceId"`
}

// ReceiveDBInstanceResponse is the response struct for api ReceiveDBInstance
type ReceiveDBInstanceResponse struct {
	*responses.BaseResponse
	GuardDBInstanceId string `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
	RequestId         string `json:"RequestId" xml:"RequestId"`
}

// CreateReceiveDBInstanceRequest creates a request to invoke ReceiveDBInstance API
func CreateReceiveDBInstanceRequest() (request *ReceiveDBInstanceRequest) {
	request = &ReceiveDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "ReceiveDBInstance", "", "")
	request.Method = requests.POST
	return
}

// CreateReceiveDBInstanceResponse creates a response to parse from ReceiveDBInstance response
func CreateReceiveDBInstanceResponse() (response *ReceiveDBInstanceResponse) {
	response = &ReceiveDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}