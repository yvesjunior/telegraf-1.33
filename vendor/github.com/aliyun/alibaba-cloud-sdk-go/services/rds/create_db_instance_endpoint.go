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

// CreateDBInstanceEndpoint invokes the rds.CreateDBInstanceEndpoint API synchronously
func (client *Client) CreateDBInstanceEndpoint(request *CreateDBInstanceEndpointRequest) (response *CreateDBInstanceEndpointResponse, err error) {
	response = CreateCreateDBInstanceEndpointResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDBInstanceEndpointWithChan invokes the rds.CreateDBInstanceEndpoint API asynchronously
func (client *Client) CreateDBInstanceEndpointWithChan(request *CreateDBInstanceEndpointRequest) (<-chan *CreateDBInstanceEndpointResponse, <-chan error) {
	responseChan := make(chan *CreateDBInstanceEndpointResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDBInstanceEndpoint(request)
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

// CreateDBInstanceEndpointWithCallback invokes the rds.CreateDBInstanceEndpoint API asynchronously
func (client *Client) CreateDBInstanceEndpointWithCallback(request *CreateDBInstanceEndpointRequest, callback func(response *CreateDBInstanceEndpointResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDBInstanceEndpointResponse
		var err error
		defer close(result)
		response, err = client.CreateDBInstanceEndpoint(request)
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

// CreateDBInstanceEndpointRequest is the request struct for api CreateDBInstanceEndpoint
type CreateDBInstanceEndpointRequest struct {
	*requests.RpcRequest
	ConnectionStringPrefix        string           `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerId               requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken                   string           `position:"Query" name:"ClientToken"`
	ResourceGroupId               string           `position:"Query" name:"ResourceGroupId"`
	DBInstanceId                  string           `position:"Query" name:"DBInstanceId"`
	DBInstanceEndpointDescription string           `position:"Query" name:"DBInstanceEndpointDescription"`
	DBInstanceEndpointType        string           `position:"Query" name:"DBInstanceEndpointType"`
	NodeItems                     string           `position:"Query" name:"NodeItems"`
	VSwitchId                     string           `position:"Query" name:"VSwitchId"`
	PrivateIpAddress              string           `position:"Query" name:"PrivateIpAddress"`
	Port                          string           `position:"Query" name:"Port"`
	VpcId                         string           `position:"Query" name:"VpcId"`
}

// CreateDBInstanceEndpointResponse is the response struct for api CreateDBInstanceEndpoint
type CreateDBInstanceEndpointResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateDBInstanceEndpointRequest creates a request to invoke CreateDBInstanceEndpoint API
func CreateCreateDBInstanceEndpointRequest() (request *CreateDBInstanceEndpointRequest) {
	request = &CreateDBInstanceEndpointRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "CreateDBInstanceEndpoint", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDBInstanceEndpointResponse creates a response to parse from CreateDBInstanceEndpoint response
func CreateCreateDBInstanceEndpointResponse() (response *CreateDBInstanceEndpointResponse) {
	response = &CreateDBInstanceEndpointResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}