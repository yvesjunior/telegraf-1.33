package cms

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

// AlarmInDescribeAlertLogList is a nested struct in cms response
type AlarmInDescribeAlertLogList struct {
	MetricName          string            `json:"MetricName" xml:"MetricName"`
	EventName           string            `json:"EventName" xml:"EventName"`
	Product             string            `json:"Product" xml:"Product"`
	BlackListUUID       string            `json:"BlackListUUID" xml:"BlackListUUID"`
	Message             string            `json:"Message" xml:"Message"`
	Namespace           string            `json:"Namespace" xml:"Namespace"`
	LevelChange         string            `json:"LevelChange" xml:"LevelChange"`
	InstanceId          string            `json:"InstanceId" xml:"InstanceId"`
	RuleName            string            `json:"RuleName" xml:"RuleName"`
	RuleId              string            `json:"RuleId" xml:"RuleId"`
	BlackListName       string            `json:"BlackListName" xml:"BlackListName"`
	GroupName           string            `json:"GroupName" xml:"GroupName"`
	GroupId             string            `json:"GroupId" xml:"GroupId"`
	AlertTime           string            `json:"AlertTime" xml:"AlertTime"`
	InstanceName        string            `json:"InstanceName" xml:"InstanceName"`
	BlackListDetail     string            `json:"BlackListDetail" xml:"BlackListDetail"`
	Level               string            `json:"Level" xml:"Level"`
	SendStatus          string            `json:"SendStatus" xml:"SendStatus"`
	LogId               string            `json:"LogId" xml:"LogId"`
	DingdingWebhookList []string          `json:"DingdingWebhookList" xml:"DingdingWebhookList"`
	ContactOnCallList   []string          `json:"ContactOnCallList" xml:"ContactOnCallList"`
	ContactMailList     []string          `json:"ContactMailList" xml:"ContactMailList"`
	ContactGroups       []string          `json:"ContactGroups" xml:"ContactGroups"`
	ContactALIIWWList   []string          `json:"ContactALIIWWList" xml:"ContactALIIWWList"`
	ContactSMSList      []string          `json:"ContactSMSList" xml:"ContactSMSList"`
	ContactDingList     []string          `json:"ContactDingList" xml:"ContactDingList"`
	SendDetail          SendDetail        `json:"SendDetail" xml:"SendDetail"`
	Escalation          Escalation        `json:"Escalation" xml:"Escalation"`
	ExtendedInfo        []ExtInfo         `json:"ExtendedInfo" xml:"ExtendedInfo"`
	Dimensions          []DimensionsItem  `json:"Dimensions" xml:"Dimensions"`
	WebhookList         []WebhookListItem `json:"WebhookList" xml:"WebhookList"`
	SendResultList      []SendResult      `json:"SendResultList" xml:"SendResultList"`
}