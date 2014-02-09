package emrysai

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getRightPassword(psd []byte) []byte {
	//做密码移动的算法
	logicStr := hex.EncodeToString(psd)
	if len(logicStr) != 32 {
		fmt.Println(len(logicStr))
		return nil
	}
	var lastFiveData = make([]byte, 5)
	tempPsd := []byte(logicStr)
	copy(lastFiveData, tempPsd[27:])
	for i := 0; i < 5; i++ {
		tempPsd[32-5+i] = tempPsd[i]
	}
	for i := 0; i < 5; i++ {
		tempPsd[i] = lastFiveData[i]
	}
	return tempPsd
}
func GetMD5Data(str string) []byte {
	pwd := md5.New()
	pwd.Write([]byte(str))
	value := pwd.Sum(nil)
	return getRightPassword(value)
}
