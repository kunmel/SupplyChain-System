package routers
import (
	"../lib"
	"../utils"
	"encoding/json"
	"fmt"

	//"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"strconv"
)
func StoreUser (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var user lib.User
	user.ID=args[0]
	user.Account=args[1]
	user.Password=args[2]
	user.Identity, _ =strconv.Atoi(args[3])
	user.Name=args[4]

	if err:=utils.WriteLedger(user,stub,lib.UserKey,[]string{user.ID});err!=nil{
		return shim.Error(fmt.Sprintf("写入User失败%s", err))
	}
	return shim.Success(nil)
}

func QueryUser(stub shim.ChaincodeStubInterface) pb.Response {
	var userList []lib.User
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.UserKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var user lib.User
			err := json.Unmarshal(v, &user)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryUser-反序列化出错: %s", err))
			}
			userList = append(userList, user)
		}
	}
	userListByte, err := json.Marshal(userList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryUser-序列化出错: %s", err))
	}
	return shim.Success(userListByte)
}

func QueryUserByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.UserKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(results[0])
}
