package routers

import (
	"fmt"
	pb "github.com/hyperledger/fabric/protos/peer"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/lib"
	"github.com/togettoyou/blockchain-real-estate/chaincode/blockchain-real-estate/utils"
	"encoding/json"
	"strconv"
)
func StoreProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	var product lib.Product
	product.GoodsID =args[0]
	product.Name =args[1]
	product.Goodtype, _ =strconv.Atoi(args[2])
	product.Amount, _ =strconv.Atoi(args[3])
	product.Price, _ =strconv.Atoi(args[4])
	product.Standards =args[5]
	product.Material =args[6]
	product.Workmanship =args[7]
	product.Supplier =args[8]
	product.SupplierID =args[9]
	product.Adddate =args[10]
	product.Tel =args[11]
	product.Shippingarea =args[12]
	product.Remark =args[13]

	if err:=utils.WriteLedger(product,stub,lib.ProductKey,[]string{product.GoodsID});err!=nil{
		return shim.Error(fmt.Sprintf("写入product失败%s", err))
	}

	return shim.Success(nil)
}


func UpdateProductByAmount (stub shim.ChaincodeStubInterface, args []string) pb.Response {
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProductKey,[]string{args[0]})
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	var product lib.Product
	_=json.Unmarshal(results[0],&product)

	product.Amount, _ =strconv.Atoi(args[1])

	if err:=utils.WriteLedger(product,stub,lib.ProductKey,[]string{product.GoodsID});err!=nil{
		return shim.Error(fmt.Sprintf("写入product失败%s", err))
	}

	return shim.Success(nil)
}

func QueryProduct(stub shim.ChaincodeStubInterface ) pb.Response {
	var productList []lib.Product
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProductKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var product lib.Product
			err := json.Unmarshal(v, &product)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryProduct-反序列化出错: %s", err))
			}
			productList = append(productList, product)
		}
	}
	productListByte, err := json.Marshal(productList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryProduct-序列化出错: %s", err))
	}
	return shim.Success(productListByte)
}

func QueryProductByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProductKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	return shim.Success(results[0])
}

func QueryProductByGoodType (stub shim.ChaincodeStubInterface ,args []string ) pb.Response{
	if len(args)!=1 {
		return shim.Error(fmt.Sprintf("参数个数!=1"))
	}
	var productList []lib.Product
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProductKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var product lib.Product
			err := json.Unmarshal(v, &product)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryProductByGoodType-反序列化出错: %s", err))
			}
			goodType,_:=strconv.Atoi(args[0])
			if product.Goodtype== goodType{
				productList = append(productList, product)
			}
		}
	}
	productListByte, err := json.Marshal(productList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryProductByGoodType-序列化出错: %s", err))
	}
	return shim.Success(productListByte)

}

func QueryProductByMaterial(stub shim.ChaincodeStubInterface ,args []string ) pb.Response{
	if len(args)!=1 {
		return shim.Error(fmt.Sprintf("参数个数!=1"))
	}
	var productList []lib.Product
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProductKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var product lib.Product
			err := json.Unmarshal(v, &product)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryProductByMaterial-反序列化出错: %s", err))
			}
			if product.Material== args[0]{
				productList = append(productList, product)
			}
		}
	}
	productListByte, err := json.Marshal(productList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryProductByMaterial-序列化出错: %s", err))
	}
	return shim.Success(productListByte)

}

func QueryProductByWorkmanship(stub shim.ChaincodeStubInterface ,args []string ) pb.Response{
	if len(args)!=1 {
		return shim.Error(fmt.Sprintf("参数个数!=1"))
	}
	var productList []lib.Product
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProductKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var product lib.Product
			err := json.Unmarshal(v, &product)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryProductByWorkmanship-反序列化出错: %s", err))
			}
			if product.Workmanship== args[0]{
				productList = append(productList, product)
			}
		}
	}
	productListByte, err := json.Marshal(productList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryProductByWorkmanship-序列化出错: %s", err))
	}
	return shim.Success(productListByte)

}

func QueryProductBySupplier(stub shim.ChaincodeStubInterface ,args []string ) pb.Response{
	if len(args)!=1 {
		return shim.Error(fmt.Sprintf("参数个数!=1"))
	}
	var productList []lib.Product
	results, err := utils.GetStateByPartialCompositeKeys2(stub, lib.ProductKey, nil)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var product lib.Product
			err := json.Unmarshal(v, &product)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryProductBySupplier-反序列化出错: %s", err))
			}
			if product.Supplier== args[0]{
				productList = append(productList, product)
			}
		}
	}
	productListByte, err := json.Marshal(productList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryProductBySupplier-序列化出错: %s", err))
	}
	return shim.Success(productListByte)
}

func DeleteProduct(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	utils.DelLedger(stub,lib.ProductKey,args)
	return shim.Success(nil)
}
