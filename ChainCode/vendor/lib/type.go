package lib

const (
	OrderKey="order-key"
	ProductKey="product-key"
	FinancingKey="financing-key"
	UserKey="uesr-key"
)

type Order struct {
	ID string//订单ID
	Buyer string//发起企业
	Seller string//供货商
	BuyerID string//企业ID
	SellerID string//供应商ID
	Address string//收货地址
	GoodsID string//对应货品ID
	Goodsname string//货品名称
	Goodsnum int //货品数量
	Orderamount int//订单总金额
	Createtime string //创建时间
	Deadline string //截止日期
	Remark string //备注
	Transferstatus int //物流状态 0：未发货， 1：已发货，2：已送达
	TransferID string //物流ID
	Status int //订单状态 0：待确认/融资， 1：已确认，2：已完成，3：供应商拒绝，4：融资被拒绝
}

type Product struct {
	GoodsID string //货品ID
	Name string//对应货品名称
	Goodtype int//货品类型 1：钢材 / 2：木方 / 3：烧结砖
	Amount int//库存
	Price int//价格
	Standards string//规格
	Material string//材料
	Workmanship string//制作工艺
	Supplier string//供应商
	SupplierID string
	Adddate string//上架时间
	Tel string//联系方式
	Shippingarea string//发货地区
	Remark string//备注
}
type User struct {
	ID string//ID(用于显示该用户)
	Account string//账号（唯一，登录时用）
	Password string//密码
	Identity int//身份(1:企业，2:供应商，3:银行)
	Name string//具体名字
}

type Financing struct {
	Supplier string //申请人
	Amount int//申请金额
	Period string//申请日期
	Desc string//申请描述
	OrderID string//对应订单号
	Status int//状态("0":待审批 "1":审批通过 "2":审批拒绝)
	Feedback string//融资反馈
}
