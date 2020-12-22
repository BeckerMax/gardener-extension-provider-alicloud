package resourcemanager

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

// ListHandshakesForResourceDirectory invokes the resourcemanager.ListHandshakesForResourceDirectory API synchronously
// api document: https://help.aliyun.com/api/resourcemanager/listhandshakesforresourcedirectory.html
func (client *Client) ListHandshakesForResourceDirectory(request *ListHandshakesForResourceDirectoryRequest) (response *ListHandshakesForResourceDirectoryResponse, err error) {
	response = CreateListHandshakesForResourceDirectoryResponse()
	err = client.DoAction(request, response)
	return
}

// ListHandshakesForResourceDirectoryWithChan invokes the resourcemanager.ListHandshakesForResourceDirectory API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/listhandshakesforresourcedirectory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListHandshakesForResourceDirectoryWithChan(request *ListHandshakesForResourceDirectoryRequest) (<-chan *ListHandshakesForResourceDirectoryResponse, <-chan error) {
	responseChan := make(chan *ListHandshakesForResourceDirectoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListHandshakesForResourceDirectory(request)
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

// ListHandshakesForResourceDirectoryWithCallback invokes the resourcemanager.ListHandshakesForResourceDirectory API asynchronously
// api document: https://help.aliyun.com/api/resourcemanager/listhandshakesforresourcedirectory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListHandshakesForResourceDirectoryWithCallback(request *ListHandshakesForResourceDirectoryRequest, callback func(response *ListHandshakesForResourceDirectoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListHandshakesForResourceDirectoryResponse
		var err error
		defer close(result)
		response, err = client.ListHandshakesForResourceDirectory(request)
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

// ListHandshakesForResourceDirectoryRequest is the request struct for api ListHandshakesForResourceDirectory
type ListHandshakesForResourceDirectoryRequest struct {
	*requests.RpcRequest
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// ListHandshakesForResourceDirectoryResponse is the response struct for api ListHandshakesForResourceDirectory
type ListHandshakesForResourceDirectoryResponse struct {
	*responses.BaseResponse
	RequestId  string                                         `json:"RequestId" xml:"RequestId"`
	PageNumber int                                            `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                                            `json:"PageSize" xml:"PageSize"`
	TotalCount int                                            `json:"TotalCount" xml:"TotalCount"`
	Handshakes HandshakesInListHandshakesForResourceDirectory `json:"Handshakes" xml:"Handshakes"`
}

// CreateListHandshakesForResourceDirectoryRequest creates a request to invoke ListHandshakesForResourceDirectory API
func CreateListHandshakesForResourceDirectoryRequest() (request *ListHandshakesForResourceDirectoryRequest) {
	request = &ListHandshakesForResourceDirectoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ResourceManager", "2020-03-31", "ListHandshakesForResourceDirectory", "resourcemanager", "openAPI")
	return
}

// CreateListHandshakesForResourceDirectoryResponse creates a response to parse from ListHandshakesForResourceDirectory response
func CreateListHandshakesForResourceDirectoryResponse() (response *ListHandshakesForResourceDirectoryResponse) {
	response = &ListHandshakesForResourceDirectoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
