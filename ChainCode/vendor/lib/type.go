package lib
type Order struct {//订单结构类型
	Buyer Account //企业
	Seller Account //供货商
	GoodName Good//商品名称
	GoodSum int//货品数量
	OrderAmount int//订单总金额
	IfAccept bool//是否接受？
	CreatTime string//订单创建时间
	Deadline string//截止日期？
	Remark string//备注
}

type Good struct {//货品结构类型
	Name string //货品名称
	Standards string //规格
	Material string  //材料
	Workmanship string  //制作工艺
	Supplier Account //供应商
	AddDate string //上架时间
	Tel string  //联系方式
	ShippingArea string  //发货地区
	Remark string //备注
}
