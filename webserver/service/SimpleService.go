/**
  author: gaoshudian
*/
package service

import (
	"fmt"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
)

type ServiceSetup struct {
	ChaincodeID string
	Client      *channel.Client
}

//赋值
func (t *ServiceSetup) SetInfo(name, num string) (string, error) {
	fmt.Println("赋值:" + name)
	return "", nil
}

//查询
func (t *ServiceSetup) GetInfo(name string) (string, error) {

	return "", nil
}
