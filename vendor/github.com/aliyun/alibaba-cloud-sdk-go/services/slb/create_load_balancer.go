package slb

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

// CreateLoadBalancer invokes the slb.CreateLoadBalancer API synchronously
func (client *Client) CreateLoadBalancer(request *CreateLoadBalancerRequest) (response *CreateLoadBalancerResponse, err error) {
	response = CreateCreateLoadBalancerResponse()
	err = client.DoAction(request, response)
	return
}

// CreateLoadBalancerWithChan invokes the slb.CreateLoadBalancer API asynchronously
func (client *Client) CreateLoadBalancerWithChan(request *CreateLoadBalancerRequest) (<-chan *CreateLoadBalancerResponse, <-chan error) {
	responseChan := make(chan *CreateLoadBalancerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateLoadBalancer(request)
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

// CreateLoadBalancerWithCallback invokes the slb.CreateLoadBalancer API asynchronously
func (client *Client) CreateLoadBalancerWithCallback(request *CreateLoadBalancerRequest, callback func(response *CreateLoadBalancerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateLoadBalancerResponse
		var err error
		defer close(result)
		response, err = client.CreateLoadBalancer(request)
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

// CreateLoadBalancerRequest is the request struct for api CreateLoadBalancer
type CreateLoadBalancerRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	LoadBalancerMode     string           `position:"Query" name:"LoadBalancerMode"`
	LoadBalancerName     string           `position:"Query" name:"LoadBalancerName"`
	Address              string           `position:"Query" name:"Address"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	IsPublicAddress      string           `position:"Query" name:"IsPublicAddress"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// CreateLoadBalancerResponse is the response struct for api CreateLoadBalancer
type CreateLoadBalancerResponse struct {
	*responses.BaseResponse
	Address          string `json:"Address" xml:"Address"`
	RequestId        string `json:"RequestId" xml:"RequestId"`
	LoadBalancerName string `json:"LoadBalancerName" xml:"LoadBalancerName"`
	LoadBalancerId   string `json:"LoadBalancerId" xml:"LoadBalancerId"`
}

// CreateCreateLoadBalancerRequest creates a request to invoke CreateLoadBalancer API
func CreateCreateLoadBalancerRequest() (request *CreateLoadBalancerRequest) {
	request = &CreateLoadBalancerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Slb", "2013-02-21", "CreateLoadBalancer", "slb", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateLoadBalancerResponse creates a response to parse from CreateLoadBalancer response
func CreateCreateLoadBalancerResponse() (response *CreateLoadBalancerResponse) {
	response = &CreateLoadBalancerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}