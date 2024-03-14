package restlib

import (
	"errors"
	"fmt"
	"github.com/prakash-p-3121/restlib/model"
	"github.com/prakash-p-3121/tomllib"

	"io/ioutil"
	"sync"
	"sync/atomic"
	"unsafe"
)

func CreateMsConnectionCfg(cfgPath string) (*sync.Map, error) {
	data, err := ioutil.ReadFile(cfgPath)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return nil, err
	}
	var cfgMap map[string]interface{}
	err = tomllib.Serialize(data, &cfgMap)
	if err != nil {
		return nil, err
	}
	var concurrentMap sync.Map
	for msName, value := range cfgMap {
		valueMap, ok := value.(map[string]interface{})
		if !ok {
			return nil, err
		}
		host := valueMap["host"].(string)
		port := valueMap["port"].(uint)

		cfg := &model.MsConnectionCfg{Host: host, Port: port}
		userPtr := unsafe.Pointer(cfg)
		newCfgPtr := atomic.SwapPointer(&userPtr, unsafe.Pointer(&model.MsConnectionCfg{
			Host: cfg.Host,
			Port: cfg.Port,
		}))

		concurrentMap.Store(msName, newCfgPtr)
	}
	return &concurrentMap, nil
}

func RetrieveMsConnectionCfg(cfgMap *sync.Map, msName string) (model.MsConnectionCfg, error) {
	valuePtr, ok := cfgMap.Load(msName)
	if !ok {
		return model.MsConnectionCfg{}, errors.New("not-found")
	}
	cfgPtr, ok := valuePtr.(*model.MsConnectionCfg)
	if !ok {
		return model.MsConnectionCfg{}, errors.New("invalid-datatype")
	}
	return *cfgPtr, nil
}
