**Data Dict**

**1.BCArDeposit	    ตารางรับเงินมัดจำ**	
Field	            Data Type	            คำอธิบาย
RowOrder	        int	Identity Field
DocNo	            varchar(20)	            เลขที่ใบมัดจำ
DocDate	            datetime	            วันที่ใบมัดจำ
TaxDate	            datetime	            วันที่ลงภาษี
TaxType	            int	                    ประเภทภาษี 0=แยกนอก 1= รวมใน 2=ภาษีอัตราศูนย์
TaxNo	            varchar(20)	            เลขที่ใบมัดจำ
ArCode	            varchar(20)	            รหัสลูกค้า
ArName	            varchar(255)	        ชื่อลูกค้า
DepartCode	        varchar(10)	            รหัสแผนก
CreditDay	        int	                    เครดิตของลูกค้า
DueDate	            datetime	            วันที่ครบกำหนดเครดิต
SaleCode	        varchar(20)	            รหัสพนักงานขาย
SaleName	        varchar(255)	        ชื่อพนักงานขาย
TaxRate	            money	                อัตราภาษี 7%
IsConfirm	        int	                    สถานะอนุมัติ 0= ไม่ได้ถูกอนุมัติหรืออ้างอิง 1=ถูกอนุมัติหรืออ้างอิงไปแล้ว
MyDescription	    varchar(255)	        หมายเหตุของเอกสาร
BeforeTaxAmount	    money	                มูลค่าก่อนภาษี
TaxAmount	        money	                มูลภาษี
TotalAmount	        money	                มูลค่าทั้งหมด
SumOfWTax	        money	                ยอดภาษีหัก ณ ที่จ่าย
NetAmount	        money	                มูลค่าสุทธิ
BillBalance	        money	                ยอดคงเหลือตัดมัดจำ
OtherIncome	        money	                รายได้อื่นๆ เพื่อให้ลงบัญชีได้
OtherExpense	    money	                ค่าใช้จ่ายอื่นๆ
ExcessAmount1	    money	                เงินเกินใบเสร็จรับ
ExcessAmount2	    money	                ยอดเงินเกินที่พัดไว้นำมาหักในใบเสร็จ
ChargeAmount	    money	                ยอด Charge บัตรเครดิต
ChangeAmount	    money	                ยอดเงินทอน
RefNo	            money	                อ้างอิงเอกสารใบสั่งขาย/สั่งจอง
CurrencyCode	    money	                รหัสสกุลเงิน
ExchangeRate	    money	                อัตราแลกเปลี่ยนเงินตรา
SumCashAmount	    money	                ยอดเงินสดรับ
SumChqAmount	    money	                ยอดเช็ครับ
SumCreditAmount 	money	                ยอดบัตรเครดิต
SumBankAmount 	    money	                ยอดเงินโอน
GLFormat	        varchar(5)	            รูปแบบการเชื่อม GL
GLStartPosting	    int	
IsPostGL	        int	                    0= ยังไม่ผ่านรายการ 1= ผ่านรายการแล้ว
IsCancel	        int	                    0= ยังไม่ยกเลิก 1= ยกเลิก
IsReturnMoney	    int	                    สถานะการนำไปลดหนี้คืนเงินมัดจำ
AllocateCode	    varchar(10)	            รหัสการจัดสรร
ProjectCode	        varchar(10)	            รหัสโครงการ
BillGroup	        varchar(10)	            กลุ่มของบิล/ภาษี
RecurName	        varchar(40)	            บันทึกเป็น Recur
ConfirmCode	        varchar(10)	            รหัสผู้อนุมัติ
ConfirmDateTime	    datetime	            วันเวลาอนุมัติ
CancelCode	        varchar(10)	            รหัสผู้ยกเลิกเอกสาร
CancelDateTime	    datetime	            วันเวลายกเลิกเอกสาร
CreatorCode	        varchar(10)	            รหัสผู้สร้างเอกสาร
CreateDateTime	    datetime	            วันเวลาสร้างเอกสาร
LastEditorCode	    varchar(10)	            รหัสผู้แก้ไขเอกสารล่าสุด
LastEditDateT	    datetime	            วันเวลาแก้ไขเอกสารล่าสุด
		
**2.BCOutPutTax	    ตารางภาษีขาย**	
SaveFrom	        int	                    บันทึกข้อมูลจาก 0=เงินมัดจำจ่าย 1=ซื้อสินค้า 2=ส่งคืนสินค้า 3=เพิ่มหนี้ 4=จ่ายชำระหนี้ 5=จ่ายเงินอื่นๆ 6=GL
DocNo	            varchar(20)	            เลขที่เอกสาร
BookCode	        varchar(15)	            รหัสสมุดรายวัน
Source	            int	                    ที่มาตอนที่โอน(มาจากระบบไหน) กำหนดไว้ให้มีค่า 6
DocDate	            datetime	            วันที่เอกสาร
TaxDate	            datetime	            วันที่ภาษี
TaxNo	            varchar(20)	            เลขที่ภาษี
ArCode	            varchar(20)	            รหัสลูกค้า
ShortTaxDesc	    varchar(100)	        คำอธิบายภาษี = ขายสินค้า
TaxRate	            money	                อัตราภาษีมูลค่าเพิ่ม
Process	            int	                    เพิ่มหรือลดภาษี กำหนดให้มีค่า 1
BeforeTaxAmount	    money	                จำนวนเงินก่อนภาษี
TaxAmount	        money	                ภาษีมูลค่าเพิ่ม
TAXID	            varchar(20)	            เลขประจำตัวผู้เสียภาษีของลูกค้า
CreatorCode	        varchar(20)	            รหัสผู้สร้างเอกสาร
CreateDateTime	    datetime	            วันเวลาสร้างเอกสาร


**API**

**ค้นหาใบมัดจำตามเลขที่เอกสาร**

GET : http://localhost:8000/ardeposit?docno=DR6011-0002

**Response**

{
    "status": "success",
    "data": {
        "row_order": 60,
        "doc_no": "DR6011-0002",
        "doc_date": "2017-11-22T00:00:00Z",
        "tax_type": 1,
        "tax_no": "DR6011-0002",
        "tax_date": "2017-11-22T00:00:00Z",
        "book_code": "03",
        "tax_id": "",
        "ar_code": "AR-03799",
        "ar_name": "บริษัท เอ็กเซคคิวทีฟ เรสซิเดนซ์ ดีเวลลอปเม้นท์ จำกัด (สำนักงานใหญ่)",
        "sale_code": "03",
        "sale_name": "สุจิตรา จันทร์พงษ์",
        "depart_code": "SALE",
        "credit_day": 0,
        "due_date": "2017-11-22T00:00:00Z",
        "tax_rate": 7,
        "is_confirm": 0,
        "my_description": "รับเงินมัดจำสินค้าสั่งผลิต ชุดคานพร้อมชุดรัดแบบ ยอดรวม907,850บาท ค้างชำระ405,850บาท",
        "before_tax_amount": 186915.89,
        "tax_amount": 13084.11,
        "total_amount": 200000,
        "sum_of_w_tax": 0,
        "net_amount": 200000,
        "bill_balance": 200000,
        "other_income": 0,
        "other_expense": 0,
        "excess_amount_1": 0,
        "excess_amount_2": 0,
        "charge_amount": 0,
        "change_amount": 0,
        "ref_no": "",
        "currency_code": "",
        "exchange_rate": 1,
        "sum_cash_amount": 0,
        "sum_chq_amount": 0,
        "sum_credit_amount": 0,
        "sum_bank_amount": 200000,
        "gl_format": "D01",
        "gl_start_posting": 0,
        "is_post_gl": 0,
        "is_cancel": 0,
        "is_return_money": 0,
        "allocate_code": "",
        "project_code": "",
        "bill_group": "",
        "recur_name": "",
        "confirm_code": "",
        "confirm_date_time": "",
        "cancel_code": "",
        "cancel_date_time": "",
        "creator_code": "kae",
        "create_date_time": "2017-11-22T10:21:31Z",
        "last_editor_code": "",
        "last_edit_date_t": ""
    }
}


**ค้นหา เอกสารใบรับเงินมัดจำ ตาคำค้นหา**

GET : http://localhost:8001/ardeposits?keyword=D


**Reponse** 

{
    "status": "success",
    "data": [
        {
            "row_order": 52,
            "doc_no": "DR6004-0001",
            "doc_date": "2017-04-26T00:00:00Z",
            "tax_type": 1,
            "tax_no": "DR6004-0001",
            "tax_date": "2017-04-26T00:00:00Z",
            "book_code": "03",
            "tax_id": "",
            "ar_code": "AR-03799",
            "ar_name": "บริษัท เอ็กเซคคิวทีฟ เรสซิเดนซ์ ดีเวลลอปเม้นท์ จำกัด (สำนักงานใหญ่)",
            "sale_code": "03",
            "sale_name": "สุจิตรา จันทร์พงษ์",
            "depart_code": "",
            "credit_day": 0,
            "due_date": "2017-04-28T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 0,
            "my_description": "เงินมัดจำค่าสินค้า",
            "before_tax_amount": 46728.97,
            "tax_amount": 3271.03,
            "total_amount": 50000,
            "sum_of_w_tax": 0,
            "net_amount": 50000,
            "bill_balance": 50000,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "charge_amount": 0,
            "change_amount": 0,
            "ref_no": "",
            "currency_code": "",
            "exchange_rate": 1,
            "sum_cash_amount": 0,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 50000,
            "gl_format": "D01",
            "gl_start_posting": 0,
            "is_post_gl": 0,
            "is_cancel": 0,
            "is_return_money": 0,
            "allocate_code": "",
            "project_code": "",
            "bill_group": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "creator_code": "jaja2525",
            "create_date_time": "2017-04-28T15:02:41Z",
            "last_editor_code": "",
            "last_edit_date_t": ""
        },
        {
            "row_order": 54,
            "doc_no": "DR6004-0002",
            "doc_date": "2017-04-29T00:00:00Z",
            "tax_type": 1,
            "tax_no": "DR6004-0002",
            "tax_date": "2017-04-29T00:00:00Z",
            "book_code": "03",
            "tax_id": "",
            "ar_code": "AR-03802",
            "ar_name": "นายธนิกร  กาจันตา",
            "sale_code": "02",
            "sale_name": "นพดล กันทา",
            "depart_code": "",
            "credit_day": 0,
            "due_date": "2017-04-29T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 0,
            "my_description": "รับมัดจำค่าซื้อแบบเหล็กขนาด 10 x 1.20 ม. จำนวน 20 แผ่น",
            "before_tax_amount": 2056.07,
            "tax_amount": 143.93,
            "total_amount": 2200,
            "sum_of_w_tax": 0,
            "net_amount": 2200,
            "bill_balance": 0,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "charge_amount": 0,
            "change_amount": 0,
            "ref_no": "",
            "currency_code": "",
            "exchange_rate": 1,
            "sum_cash_amount": 2200,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 0,
            "gl_format": "D01",
            "gl_start_posting": 0,
            "is_post_gl": 0,
            "is_cancel": 0,
            "is_return_money": 0,
            "allocate_code": "",
            "project_code": "",
            "bill_group": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "creator_code": "sale",
            "create_date_time": "2017-04-29T15:05:45Z",
            "last_editor_code": "",
            "last_edit_date_t": ""
        },
        {
            "row_order": 55,
            "doc_no": "DR6004-0003",
            "doc_date": "2017-04-29T00:00:00Z",
            "tax_type": 1,
            "tax_no": "DR6004-0003",
            "tax_date": "2017-04-29T00:00:00Z",
            "book_code": "03",
            "tax_id": "",
            "ar_code": "AR-03799",
            "ar_name": "บริษัท เอ็กเซคคิวทีฟ เรสซิเดนซ์ ดีเวลลอปเม้นท์ จำกัด (สำนักงานใหญ่)",
            "sale_code": "03",
            "sale_name": "สุจิตรา จันทร์พงษ์",
            "depart_code": "",
            "credit_day": 0,
            "due_date": "2017-05-02T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 0,
            "my_description": "",
            "before_tax_amount": 88650.47,
            "tax_amount": 6205.53,
            "total_amount": 94856,
            "sum_of_w_tax": 0,
            "net_amount": 94856,
            "bill_balance": 94856,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "charge_amount": 0,
            "change_amount": 0,
            "ref_no": "",
            "currency_code": "",
            "exchange_rate": 1,
            "sum_cash_amount": 0,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 94856,
            "gl_format": "D01",
            "gl_start_posting": 0,
            "is_post_gl": 0,
            "is_cancel": 0,
            "is_return_money": 0,
            "allocate_code": "",
            "project_code": "",
            "bill_group": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "creator_code": "tawon",
            "create_date_time": "2017-05-02T08:51:50Z",
            "last_editor_code": "",
            "last_edit_date_t": ""
        },
        {
            "row_order": 59,
            "doc_no": "DR6011-0001",
            "doc_date": "2017-11-07T00:00:00Z",
            "tax_type": 1,
            "tax_no": "DR6011-0001",
            "tax_date": "2017-11-07T00:00:00Z",
            "book_code": "03",
            "tax_id": "",
            "ar_code": "AR-03799",
            "ar_name": "บริษัท เอ็กเซคคิวทีฟ เรสซิเดนซ์ ดีเวลลอปเม้นท์ จำกัด (สำนักงานใหญ่)",
            "sale_code": "02",
            "sale_name": "นพดล กันทา",
            "depart_code": "SALE",
            "credit_day": 0,
            "due_date": "2017-11-07T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 0,
            "my_description": "รับเงินมัดจำค่าสั่งซื้อสินค้า 1.ชุดยกบล็อกคอนกรีต 2 ชุด 2.ชุดแบบพิเศษบ้านน็อคดาวน์ 4 ชุด ค้างชำระ103,850บาท",
            "before_tax_amount": 280373.83,
            "tax_amount": 19626.17,
            "total_amount": 300000,
            "sum_of_w_tax": 0,
            "net_amount": 300000,
            "bill_balance": 300000,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "charge_amount": 0,
            "change_amount": 0,
            "ref_no": "",
            "currency_code": "",
            "exchange_rate": 1,
            "sum_cash_amount": 0,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 300000,
            "gl_format": "D01",
            "gl_start_posting": 0,
            "is_post_gl": 0,
            "is_cancel": 0,
            "is_return_money": 0,
            "allocate_code": "",
            "project_code": "",
            "bill_group": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "creator_code": "kae",
            "create_date_time": "2017-11-07T09:17:11Z",
            "last_editor_code": "",
            "last_edit_date_t": ""
        },
        {
            "row_order": 60,
            "doc_no": "DR6011-0002",
            "doc_date": "2017-11-22T00:00:00Z",
            "tax_type": 1,
            "tax_no": "DR6011-0002",
            "tax_date": "2017-11-22T00:00:00Z",
            "book_code": "03",
            "tax_id": "",
            "ar_code": "AR-03799",
            "ar_name": "บริษัท เอ็กเซคคิวทีฟ เรสซิเดนซ์ ดีเวลลอปเม้นท์ จำกัด (สำนักงานใหญ่)",
            "sale_code": "03",
            "sale_name": "สุจิตรา จันทร์พงษ์",
            "depart_code": "SALE",
            "credit_day": 0,
            "due_date": "2017-11-22T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 0,
            "my_description": "รับเงินมัดจำสินค้าสั่งผลิต ชุดคานพร้อมชุดรัดแบบ ยอดรวม907,850บาท ค้างชำระ405,850บาท",
            "before_tax_amount": 186915.89,
            "tax_amount": 13084.11,
            "total_amount": 200000,
            "sum_of_w_tax": 0,
            "net_amount": 200000,
            "bill_balance": 200000,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "charge_amount": 0,
            "change_amount": 0,
            "ref_no": "",
            "currency_code": "",
            "exchange_rate": 1,
            "sum_cash_amount": 0,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 200000,
            "gl_format": "D01",
            "gl_start_posting": 0,
            "is_post_gl": 0,
            "is_cancel": 0,
            "is_return_money": 0,
            "allocate_code": "",
            "project_code": "",
            "bill_group": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "creator_code": "kae",
            "create_date_time": "2017-11-22T10:21:31Z",
            "last_editor_code": "",
            "last_edit_date_t": ""
        },
        {
            "row_order": 77,
            "doc_no": "S01-DP6102-0001",
            "doc_date": "2018-02-13T00:00:00Z",
            "tax_type": 1,
            "tax_no": "S01-DP6102-0001",
            "tax_date": "2018-02-13T00:00:00Z",
            "book_code": "6",
            "tax_id": "5557231229",
            "ar_code": "45040",
            "ar_name": "",
            "sale_code": "56163",
            "sale_name": "",
            "depart_code": "",
            "credit_day": 1,
            "due_date": "2018-02-13T00:00:00Z",
            "tax_rate": 7,
            "is_confirm": 0,
            "my_description": "Test",
            "before_tax_amount": 93,
            "tax_amount": 7,
            "total_amount": 100,
            "sum_of_w_tax": 0,
            "net_amount": 100,
            "bill_balance": 100,
            "other_income": 0,
            "other_expense": 0,
            "excess_amount_1": 0,
            "excess_amount_2": 0,
            "charge_amount": 0,
            "change_amount": 0,
            "ref_no": "Ref",
            "currency_code": "",
            "exchange_rate": 1,
            "sum_cash_amount": 100,
            "sum_chq_amount": 0,
            "sum_credit_amount": 0,
            "sum_bank_amount": 0,
            "gl_format": "",
            "gl_start_posting": 0,
            "is_post_gl": 0,
            "is_cancel": 0,
            "is_return_money": 0,
            "allocate_code": "",
            "project_code": "",
            "bill_group": "",
            "recur_name": "",
            "confirm_code": "",
            "confirm_date_time": "",
            "cancel_code": "",
            "cancel_date_time": "",
            "creator_code": "moo",
            "create_date_time": "2018-02-13T16:18:19.053Z",
            "last_editor_code": "",
            "last_edit_date_t": ""
        }
    ]
}

**บันทึก ใบรับเงินมัดจำ**

POST : http://localhost:8000/ardeposit

payload = {
	"doc_no":"S01-DP6102-0001",
	"doc_date":"2018/02/13",
	"tax_date":"2018/02/13",
	"tax_type":1,
	"tax_no":"S01-DP6102-0001",
	"book_code":"6",
	"tax_id":"5557231229",
	"ar_code":"45040",
	"ar_name":"",
	"depart_code":"",
	"credit_day":1,
	"due_date":"2018/02/13",
	"sale_code":"56163",
	"sale_name":"",
	"tax_rate":7,
	"is_confirm":0,
	"my_description":"",
	"before_tax_amount":93,
	"tax_amount":7,
	"total_amount":100,
	"sum_of_w_tax":0,
	"net_amount":100,
	"bill_balance":100,
	"other_income":0,
	"other_expense":0,
	"ref_no":"Ref",
	"exchange_rate":1.0,
	"sum_cash_amount":100,
	"creator_code":"moo"
}



**เปลี่ยนแปลง ใบรับเงินมัดจำ**

PUT : http://localhost:8000/ardeposit

payload = {
	"doc_no":"S01-DP6102-0001",
	"doc_date":"2018/02/13",
	"tax_date":"2018/02/13",
	"tax_type":1,
	"tax_no":"S01-DP6102-0001",
	"book_code":"6",
	"tax_id":"5557231229",
	"ar_code":"45040",
	"ar_name":"",
	"depart_code":"",
	"credit_day":1,
	"due_date":"2018/02/13",
	"sale_code":"56163",
	"sale_name":"",
	"tax_rate":7,
	"is_confirm":0,
	"my_description":"",
	"before_tax_amount":93,
	"tax_amount":7,
	"total_amount":100,
	"sum_of_w_tax":0,
	"net_amount":100,
	"bill_balance":100,
	"other_income":0,
	"other_expense":0,
	"ref_no":"Ref",
	"exchange_rate":1.0,
	"sum_cash_amount":100,
	"creator_code":"moo"
}