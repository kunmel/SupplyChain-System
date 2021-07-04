package routers

import (
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"strconv"
	"encoding/json"
)
func StoreFinance (stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var financing lib.Financing
	financing.Supplier =args[0]
	financing.Amount, _ =strconv.Atoi(args[1])
	financing.Period =args[2]
	financing.Desc =args[3]
	financing.OrderID =args[4]
	financing.Status, _ =strconv.Atoi(args[5])
	financing.Feedback =args[6]

	if err:=utils.WriteLedger(financing,stub,lib.FinancingKey,[]string{financing.OrderID,financing.Period});err!=nil{
		return shim.Error(fmt.Sprintf("写入Financing失败%s", err))
	}

	return shim.Success(nil)
}

func UpdateFinanceStatus (stub shim.ChaincodeStubInterface, args []string) pb.Response{//args[0]:orderid,args[1]:period,arg[2]:status,arg[3]:Feedback
 results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.FinancingKey,[]string{args[0],args[1]})
 if err != nil {
  return shim.Error(fmt.Sprintf("%s", err))
 }
 var financing lib.Financing
 _=json.Unmarshal(results[0],&financing)

 financing.Status, _ =strconv.Atoi(args[2])
 financing.Feedback =args[3]

 if err:=utils.WriteLedger(financing,stub,lib.FinancingKey,[]string{financing.OrderID,financing.Period,financing.Feedback});err!=nil{
  return shim.Error(fmt.Sprintf("写入Financing失败%s", err))
 }

 return shim.Success(nil)
}

func QueryFinance(stub shim.ChaincodeStubInterface) pb.Response{
	var financingList []lib.Financing
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.FinancingKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var financing lib.Financing
			err := json.Unmarshal(v, &financing)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryFinance-反序列化出错: %s", err))
			}
			financingList = append(financingList, financing)
		}
	}
	financingListByte, err := json.Marshal(financingList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryFinance-序列化出错: %s", err))
	}
	return shim.Success(financingListByte)
}

func QueryFinanceBySupplier(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	 if len(args)!=1 {
	  return shim.Error(fmt.Sprintf("参数个数!=1"))
	 }
	 var financingList []lib.Financing
	 results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.FinancingKey, nil)
	 if err != nil {
	  return shim.Error(fmt.Sprintf("%s", err))
	 }
	 for _, v := range results {
	  if v != nil {
	   var financing lib.Financing
	   err := json.Unmarshal(v, &financing)
	   if err != nil {
	    return shim.Error(fmt.Sprintf("QueryFinanceByStatus-反序列化出错: %s", err))
	   }
	   if financing.Supplier==args[0] {
	    financingList = append(financingList, financing)
	   }
	  }
	 }
	 financingListByte, err := json.Marshal(financingList)
	 if err != nil {
	  return shim.Error(fmt.Sprintf("QueryFinanceByStatus-序列化出错: %s", err))
	 }
	 return shim.Success(financingListByte)
}

func QueryFinanceByStatus(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args)!=1 {
		return shim.Error(fmt.Sprintf("参数个数!=1"))
	}
	var financingList []lib.Financing
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.FinancingKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var financing lib.Financing
			err := json.Unmarshal(v, &financing)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryFinanceByStatus-反序列化出错: %s", err))
			}
			status,_:=strconv.Atoi(args[0])
			if financing.Status==status {
				financingList = append(financingList, financing)
			}
		}
	}
	financingListByte, err := json.Marshal(financingList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryFinanceByStatus-序列化出错: %s", err))
	}
	return shim.Success(financingListByte)
}


func QueryFinanceByOrderID(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	if len(args)!=1 {
		return shim.Error(fmt.Sprintf("参数个数!=1"))
	}
	var financingList []lib.Financing
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.FinancingKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var financing lib.Financing
			err := json.Unmarshal(v, &financing)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryFinanceByOrderID-反序列化出错: %s", err))
			}
			if financing.OrderID==args[0] {
				financingList = append(financingList, financing)
			}
		}
	}
	financingListByte, err := json.Marshal(financingList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryFinanceByOrderID-序列化出错: %s", err))
	}
	return shim.Success(financingListByte)
}



