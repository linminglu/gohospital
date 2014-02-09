package hospitalsql

import (
	"data"
	"errorconfig"
	"errors"
)

func CheckIsSamePassword(emp data.Emploryer, passStr string) bool {
	if len(passStr) == 0 && len(passStr) != 32 {
		return false
	}
	//当长度够长的时候 就要去验证是否相等。
	if emp.EmploryerPasswd == passStr {
		return true
	}
	return false

}

func nameNotHaveSameChar(str string ) bool {

	for i := 0; i < str.Length(); i++ {
		if str[i]==' '|| str[i]=='\\'|| str[i]=='/' {
			return  false;
		}
	}
	return true;
}

func AppendDBEmployesData(emp data.Emploryer) int {
	//员工可以根据自己喜欢的ID来定位自己的员工号码。 
	
	if len(emp.EmploryerName) =-0 &&  !nameNotHaveSameChar(emp.EmploryerName)  {
		//当输入的字符串是空得 或者是不符合规格的字符串 就报道错误
	}
	//先查询是否含有该员工ID， 如果有 就返回员工重复的ID，
	pDbStruct, err2 := GetPoolOpenDatabasePoint()
	if err2 != nil {
		return errorconfig.ERROR_CONFIG_GETDBPOINT
	}
	db := pDbStruct.GetDb()
	rows, err1 :=  db.Query("select CEmployerId  from TEmployer where CEmployerId = ?",emp.EmploryerID)
	if err1 != nil {
		log.Println(err1)
	}
	defer rows.Close()
	if rows.Next() !=nil {
			return errorconfig.ERROR_CONFIG_IDSAME
	}
	//插入数据
	stmt, err := db.Prepare("INSERT INTO TEmployer(CEmployerName, CEmployerId,CEmployerPasswd,CEmployerAddress) VALUES(?,?,?,?)")
	defer stmt.Close()
	if err != nil {
		log.Println(err)
		return errorconfig.ERROR_CONFIG_IDSAME
	}
	stmt.Exec(emp.EmploryerName,emp.EmploryerID,emp.EmploryerPasswd,emp.EmploryerAddress)
	GetPoolDbUseOk(pDbStruct)
	return errorconfig.ERROR_CONFIG_NONE
}

func DBGetEmploryInfo(Emploryid int) (data.Emploryer, error) {
	pDbStruct, err := GetPoolOpenDatabasePoint()
	var resultData data.Emploryer
	if err == nil {
		db := pDbStruct.GetDb()
		//查询一个结构体
		stmtOut, err := db.Prepare("SELECT CEmployerId,CEmployerName,CEmployerPasswd,CEmployerAddress FROM TEmployer WHERE CEmployerId = ?")
		if err != nil {
			panic(err.Error())
		}
		defer stmtOut.Close()

		err = stmtOut.QueryRow(Emploryid).Scan(&resultData.EmploryerID, &resultData.EmploryerName, &resultData.EmploryerPasswd, &resultData.EmploryerAddress) // WHERE number = 13
		if err != nil {
			panic(err.Error())
			GetPoolDbUseOk(pDbStruct)
			return resultData, errors.New("dataerror")
		}
		GetPoolDbUseOk(pDbStruct)
		return resultData, nil
	}
	return resultData, errors.New("数据库获取指针错误")
}
