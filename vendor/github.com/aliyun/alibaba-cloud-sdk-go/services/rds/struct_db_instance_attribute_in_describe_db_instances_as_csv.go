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

// DBInstanceAttributeInDescribeDBInstancesAsCsv is a nested struct in rds response
type DBInstanceAttributeInDescribeDBInstancesAsCsv struct {
	VpcId                       string                               `json:"VpcId" xml:"VpcId"`
	CreationTime                string                               `json:"CreationTime" xml:"CreationTime"`
	TempDBInstanceId            string                               `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
	SupportUpgradeAccountType   string                               `json:"SupportUpgradeAccountType" xml:"SupportUpgradeAccountType"`
	IncrementSourceDBInstanceId string                               `json:"IncrementSourceDBInstanceId" xml:"IncrementSourceDBInstanceId"`
	DBInstanceMemory            int64                                `json:"DBInstanceMemory" xml:"DBInstanceMemory"`
	MaintainTime                string                               `json:"MaintainTime" xml:"MaintainTime"`
	PayType                     string                               `json:"PayType" xml:"PayType"`
	Tags                        string                               `json:"Tags" xml:"Tags"`
	AvailabilityValue           string                               `json:"AvailabilityValue" xml:"AvailabilityValue"`
	ReadDelayTime               string                               `json:"ReadDelayTime" xml:"ReadDelayTime"`
	ConnectionMode              string                               `json:"ConnectionMode" xml:"ConnectionMode"`
	Port                        string                               `json:"Port" xml:"Port"`
	AccountType                 string                               `json:"AccountType" xml:"AccountType"`
	LockMode                    string                               `json:"LockMode" xml:"LockMode"`
	EngineVersion               string                               `json:"EngineVersion" xml:"EngineVersion"`
	MaxIOPS                     int                                  `json:"MaxIOPS" xml:"MaxIOPS"`
	ConnectionString            string                               `json:"ConnectionString" xml:"ConnectionString"`
	InstanceNetworkType         string                               `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	SecurityIPList              string                               `json:"SecurityIPList" xml:"SecurityIPList"`
	MasterInstanceId            string                               `json:"MasterInstanceId" xml:"MasterInstanceId"`
	DBInstanceClassType         string                               `json:"DBInstanceClassType" xml:"DBInstanceClassType"`
	DBInstanceDescription       string                               `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	DBInstanceCPU               string                               `json:"DBInstanceCPU" xml:"DBInstanceCPU"`
	ExpireTime                  string                               `json:"ExpireTime" xml:"ExpireTime"`
	DBInstanceNetType           string                               `json:"DBInstanceNetType" xml:"DBInstanceNetType"`
	DBInstanceType              string                               `json:"DBInstanceType" xml:"DBInstanceType"`
	AccountMaxQuantity          int                                  `json:"AccountMaxQuantity" xml:"AccountMaxQuantity"`
	LockReason                  string                               `json:"LockReason" xml:"LockReason"`
	DBInstanceStatus            string                               `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	DBMaxQuantity               int                                  `json:"DBMaxQuantity" xml:"DBMaxQuantity"`
	GuardDBInstanceId           string                               `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
	RegionId                    string                               `json:"RegionId" xml:"RegionId"`
	DBInstanceStorage           int                                  `json:"DBInstanceStorage" xml:"DBInstanceStorage"`
	VSwitchId                   string                               `json:"VSwitchId" xml:"VSwitchId"`
	ZoneId                      string                               `json:"ZoneId" xml:"ZoneId"`
	Category                    string                               `json:"Category" xml:"Category"`
	MaxConnections              int                                  `json:"MaxConnections" xml:"MaxConnections"`
	DBInstanceId                string                               `json:"DBInstanceId" xml:"DBInstanceId"`
	DBInstanceClass             string                               `json:"DBInstanceClass" xml:"DBInstanceClass"`
	Engine                      string                               `json:"Engine" xml:"Engine"`
	ExportKey                   string                               `json:"ExportKey" xml:"ExportKey"`
	SlaveZones                  SlaveZonesInDescribeDBInstancesAsCsv `json:"SlaveZones" xml:"SlaveZones"`
}