package model

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"errors"
	"database/sql"
)

type Item struct {
	Code            string  `json:"code" db:"Code"`
	Name1           string  `json:"name_1" db:"Name1"`
	Name2           string  `json:"name_2" db:"Name2"`
	CreatorCode     string  `json:"creator_code" db:"CreatorCode"`
	CreateDateTime  string  `json:"create_date_time" db:"CreateDateTime"`
	LastEditorCode  string  `json:"last_editor_code" db:"LastEditorCode"`
	LastEditDateT   string  `json:"last_edit_date_t" db:"LastEditDateT"`
	ShortName       string  `json:"short_name" db:"ShortName"`
	CategoryCode    string  `json:"category_code" db:"CategoryCode"`
	GroupCode       string  `json:"group_code" db:"GroupCode"`
	BrandCode       string  `json:"brand_code" db:"BrandCode"`
	TypeCode        string  `json:"type_code" db:"TypeCode"`
	FormatCode      string  `json:"format_code" db:"FormatCode"`
	ColorCode       string  `json:"color_code" db:"ColorCode"`
	UnitType        int     `json:"unit_type" db:"UnitType"`
	DefStkUnitCode  string  `json:"def_stk_unit_code" db:"DefStkUnitCode"`
	DefSaleUnitCode string  `json:"def_sale_unit_code" db:"DefSaleUnitCode"`
	DefBuyUnitCode  string  `json:"def_buy_unit_code" db:"DefBuyUnitCode"`
	StockType       int     `json:"stock_type" db:"StockType"`
	ItemStatus      int     `json:"item_status" db:"ItemStatus"`
	OrderPoint      float64 `json:"order_point" db:"OrderPoint"`
	StockMin        float64 `json:"stock_min" db:"StockMin"`
	StockMax        float64 `json:"stock_max" db:"StockMax"`
	StockQty        float64 `json:"stock_qty" db:"StockQty"`
	AverageCost     float64 `json:"average_cost" db:"AverageCost"`
	PicFileName1    string  `json:"pic_file_name_1" db:"PicFileName1"`
	PicFileName2    string  `json:"pic_file_name_2" db:"PicFileName2"`
	MyDescription   string  `json:"my_description" db:"MyDescription"`
	DefBuyWHCode    string  `json:"def_buy_wh_code" db:"DefBuyWHCode"`
	DefBuyShelf     string  `json:"def_buy_shelf" db:"DefBuyShelf"`
	DefSaleWHCode   string  `json:"def_sale_wh_code" db:"DefSaleWHCode"`
	DefSaleShelf    string  `json:"def_sale_shelf" db:"DefSaleShelf"`
	ActiveStatus    int     `json:"active_status" db:"ActiveStatus"`
	SalePrice1      float64 `json:"sale_price_1" db:"SalePrice1"`
	SalePrice2      float64 `json:"sale_price_2" db:"SalePrice2"`
	SalePrice3      float64 `json:"sale_price_3" db:"SalePrice3"`
	SalePrice4      float64 `json:"sale_price_4" db:"SalePrice4"`
	DefFixUnitCode  string  `json:"def_fix_unit_code" db:"DefFixUnitCode"`
}

func (itm *Item) SearchItemByCode(db *sqlx.DB, code string) error {
	sql := `set dateformat dmy     select Code,Name1,isnull(Name2,'') as Name2,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(ShortName,'') as ShortName,isnull(CategoryCode,'') as CategoryCode,isnull(GroupCode,'') as GroupCode,isnull(BrandCode,'') as BrandCode,isnull(TypeCode,'') as TypeCode,isnull(FormatCode,'') as FormatCode,isnull(ColorCode,'') as ColorCode,UnitType,isnull(DefStkUnitCode,'') as DefStkUnitCode,isnull(DefSaleUnitCode,'') as DefSaleUnitCode,isnull(DefBuyUnitCode,'') as DefBuyUnitCode,StockType,isnull(ItemStatus,0) as ItemStatus,OrderPoint,StockMin,StockMax,StockQty,AverageCost,isnull(PicFileName1,'') as PicFileName1,isnull(PicFileName2,'') as PicFileName2,isnull(MyDescription,'') as MyDescription,isnull(DefBuyWHCode,'') as DefBuyWHCode,isnull(DefBuyShelf,'') as DefBuyShelf,isnull(DefSaleWHCode,'') as DefSaleWHCode,isnull(DefSaleShelf,'') as DefSaleShelf,ActiveStatus,isnull(SalePrice1,0) as SalePrice1,isnull(SalePrice2,0) as SalePrice2,isnull(SalePrice3,0) as SalePrice3,isnull(SalePrice4,0) as SalePrice4,isnull(DefFixUnitCode,'') as DefFixUnitCode from dbo.bcitem where code = ?`
	fmt.Println("sql = ", sql, code)
	err := db.Get(itm, sql, code)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return err
	}

	return nil
}

func (itm *Item) SearchItemByKeycode(db *sqlx.DB, keyword string) (itmList []*Item, err error) {
	sql := `set dateformat dmy     select Code,Name1,isnull(Name2,'') as Name2,isnull(CreatorCode,'') as CreatorCode,isnull(CreateDateTime,'') as CreateDateTime,isnull(LastEditorCode,'') as LastEditorCode,isnull(LastEditDateT,'') as LastEditDateT,isnull(ShortName,'') as ShortName,isnull(CategoryCode,'') as CategoryCode,isnull(GroupCode,'') as GroupCode,isnull(BrandCode,'') as BrandCode,isnull(TypeCode,'') as TypeCode,isnull(FormatCode,'') as FormatCode,isnull(ColorCode,'') as ColorCode,UnitType,isnull(DefStkUnitCode,'') as DefStkUnitCode,isnull(DefSaleUnitCode,'') as DefSaleUnitCode,isnull(DefBuyUnitCode,'') as DefBuyUnitCode,StockType,isnull(ItemStatus,0) as ItemStatus,OrderPoint,StockMin,StockMax,StockQty,AverageCost,isnull(PicFileName1,'') as PicFileName1,isnull(PicFileName2,'') as PicFileName2,isnull(MyDescription,'') as MyDescription,isnull(DefBuyWHCode,'') as DefBuyWHCode,isnull(DefBuyShelf,'') as DefBuyShelf,isnull(DefSaleWHCode,'') as DefSaleWHCode,isnull(DefSaleShelf,'') as DefSaleShelf,ActiveStatus,isnull(SalePrice1,0) as SalePrice1,isnull(SalePrice2,0) as SalePrice2,isnull(SalePrice3,0) as SalePrice3,isnull(SalePrice4,0) as SalePrice4,isnull(DefFixUnitCode,'') as DefFixUnitCode from dbo.bcitem where (code  like '%'+?+'%' or name1 like '%'+?+'%' or name2 like '%'+?+'%' )`
	fmt.Println("sql = ", sql, keyword)
	err = db.Select(&itmList, sql, keyword, keyword, keyword)
	if err != nil {
		fmt.Println("error = ", err.Error())
		return nil, err
	}

	return itmList, nil
}

func (itm *Item) SaveItem(db *sqlx.DB) error {
	var check_exist int

	sqlexist := `select count(code) as check_exist from dbo.bcitem where code = ?`
	err := db.Get(&check_exist, sqlexist, itm.Code)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	switch {
	case itm.Code == "":
		return errors.New("code is null")
	case itm.Name1 == "":
		return errors.New("Name1 is null")
	case itm.DefStkUnitCode == "":
		return errors.New("unitcode is null")
	case itm.DefBuyWHCode == "":
		return errors.New("buy whcode is null")
	case itm.DefSaleWHCode == "":
		return errors.New("sale whcode is null")
	case itm.DefBuyShelf == "":
		return errors.New("buy shelf is null")
	case itm.DefSaleShelf == "":
		return errors.New("sale shelf is null")
	}

	var i sql.NullInt64
	var errint error

	switch {
	case  errint == i.Scan(itm.UnitType) && errint != nil :
		itm.UnitType = 0
	case  errint == i.Scan(itm.StockType) && errint != nil :
		itm.StockType = 0
	case  errint == i.Scan(itm.ItemStatus) && errint != nil :
		itm.StockType = 1
	case itm.DefStkUnitCode != "":
		itm.DefFixUnitCode = itm.DefStkUnitCode
	case itm.DefStkUnitCode != "" && itm.DefSaleUnitCode == "":
		itm.DefSaleUnitCode = itm.DefStkUnitCode
	case itm.DefStkUnitCode != "" && itm.DefBuyUnitCode == "":
		itm.DefBuyUnitCode = itm.DefStkUnitCode
	}

	if (check_exist == 0) {
		itm.ActiveStatus = 1

		sql := `Insert into dbo.bcitem(Code,Name1,Name2,CreatorCode,CreateDateTime,ShortName,CategoryCode,GroupCode,BrandCode,TypeCode,FormatCode,ColorCode,UnitType,DefStkUnitCode,DefSaleUnitCode,DefBuyUnitCode,StockType,ItemStatus,OrderPoint,StockMin,StockMax,StockQty,AverageCost,PicFileName1,PicFileName2,MyDescription,DefBuyWHCode,DefBuyShelf,DefSaleWHCode,DefSaleShelf,ActiveStatus,SalePrice1,SalePrice2,SalePrice3,SalePrice4,DefFixUnitCode) values(?,?,?,?,getdate(),?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
		_, err = db.Exec(sql, itm.Code,itm.Name1,itm.Name2,itm.CreatorCode,itm.ShortName,itm.CategoryCode,itm.GroupCode,itm.BrandCode,itm.TypeCode,itm.FormatCode,itm.ColorCode,itm.UnitType,itm.DefStkUnitCode,itm.DefSaleUnitCode,itm.DefBuyUnitCode,itm.StockType,itm.ItemStatus,itm.OrderPoint,itm.StockMin,itm.StockMax,itm.StockQty,itm.AverageCost,itm.PicFileName1,itm.PicFileName2,itm.MyDescription,itm.DefBuyWHCode,itm.DefBuyShelf,itm.DefSaleWHCode,itm.DefSaleShelf,itm.ActiveStatus,itm.SalePrice1,itm.SalePrice2,itm.SalePrice3,itm.SalePrice4,itm.DefFixUnitCode)
	} else {
		sql := `Update dbo.bcitem set Name1=?,Name2=?,LastEditorCode=?,LastEditDateT=getdate(),ShortName=?,CategoryCode=?,GroupCode=?,BrandCode=?,TypeCode=?,FormatCode=?,ColorCode=?,UnitType=?,DefStkUnitCode=?,DefSaleUnitCode=?,DefBuyUnitCode=?,StockType=?,ItemStatus=?,OrderPoint=?,StockMin=?,StockMax=?,StockQty=?,AverageCost=?,PicFileName1=?,PicFileName2=?,MyDescription=?,DefBuyWHCode=?,DefBuyShelf=?,DefSaleWHCode=?,DefSaleShelf=?,ActiveStatus=?,SalePrice1=?,SalePrice2=?,SalePrice3=?,SalePrice4=?,DefFixUnitCode=? where code = ?`
		_, err = db.Exec(sql, itm.Name1,itm.Name2,itm.LastEditorCode,itm.ShortName,itm.CategoryCode,itm.GroupCode,itm.BrandCode,itm.TypeCode,itm.FormatCode,itm.ColorCode,itm.UnitType,itm.DefStkUnitCode,itm.DefSaleUnitCode,itm.DefBuyUnitCode,itm.StockType,itm.ItemStatus,itm.OrderPoint,itm.StockMin,itm.StockMax,itm.StockQty,itm.AverageCost,itm.PicFileName1,itm.PicFileName2,itm.MyDescription,itm.DefBuyWHCode,itm.DefBuyShelf,itm.DefSaleWHCode,itm.DefSaleShelf,itm.ActiveStatus,itm.SalePrice1,itm.SalePrice2,itm.SalePrice3,itm.SalePrice4, itm.DefFixUnitCode,itm.Code)
	}
	if err != nil {
		fmt.Println("error = ", err.Error())
		return err
	}

	return nil
}
