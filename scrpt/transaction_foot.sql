/*
 * 描述：邯郸--商家交易表
 *
 *  transaction : 交易ID( 20190122 )
 *  tra_type	: 交易类型 : 0 锘 , 1 现金
 *  status      ：0 支付, 1 退款
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS car_transaction_310;
CREATE TABLE IF NOT EXISTS car_transaction_310(
	id		INT UNSIGNED		NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '表    ID',
	tran_id		CHAR(60)		NOT NULL		COMMENT '交 易 ID',
	user_id		CHAR(32)   	    	NOT NULL 		COMMENT '用 户 ID',
	user_phone	CHAR(11)   	    	NOT NULL 		COMMENT '用户手机号',
	user_name	VARCHAR(64)   	    	NOT NULL 		COMMENT '用户姓名',
	merchant_name   VARCHAR(128)		NOT NULL		COMMENT '商家名称',
	cashier_id	CHAR(32)		NOT NULL		COMMENT '收 银 ID',
	merchant_id	CHAR(32)   	    	NOT NULL 		COMMENT '商 家 ID',
	rate		DOUBLE(10,2)		NOT NULL		COMMENT '本单费率金额',
	amount		DOUBLE(10,2)		NOT NULL		COMMENT '交易金额',
	refund_amount	DOUBLE(10,2)		NOT NULL		COMMENT '退款金额',
	mer_balance	DOUBLE(14,2)		NOT NULL		COMMENT '商户余额',
	tran_type	INT UNSIGNED		NOT NULL		COMMENT '交易类型',
	note_test	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '备    注',
	status		INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT '当前状态',
	create_at	INT UNSIGNED		NOT NULL		COMMENT	'创建时间',
	update_at	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'更新时间',
	INDEX car_transaction_310( user_id, merchant_id )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='邯郸商家交易表';

/*
 * 描述：石家庄--商家交易表
 *
 *  transaction : 交易ID( 20190122 )
 *  tra_type	: 交易类型 : 0 锘 , 1 现金
 *  status      ：0 支付, 1 退款
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS car_transaction_311;
CREATE TABLE IF NOT EXISTS car_transaction_311(
	id		INT UNSIGNED		NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '表    ID',
	tran_id		CHAR(60)		NOT NULL		COMMENT '交 易 ID',
	user_id		CHAR(32)   	    	NOT NULL 		COMMENT '用 户 ID',
	user_phone	CHAR(11)   	    	NOT NULL 		COMMENT '用户手机号',
	user_name	VARCHAR(64)   	    	NOT NULL 		COMMENT '用户姓名',
	cashier_id	CHAR(32)		NOT NULL		COMMENT '收 银 ID',
	merchant_id	CHAR(32)   	    	NOT NULL 		COMMENT '商 家 ID',
	rate		DOUBLE(10,2)		NOT NULL		COMMENT '本单费率金额',
	amount		DOUBLE(10,2)		NOT NULL		COMMENT '交易金额',
	refund_amount	DOUBLE(10,2)		NOT NULL		COMMENT '退款金额',
	mer_balance	DOUBLE(14,2)		NOT NULL		COMMENT '商户余额',
	tran_type	INT UNSIGNED		NOT NULL		COMMENT '交易类型',
	note_test	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '备    注',
	status		INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT '当前状态',
	create_at	INT UNSIGNED		NOT NULL		COMMENT	'创建时间',
	update_at	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'更新时间',
	INDEX car_transaction_311( user_id, merchant_id )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='石家庄商家交易表';

/*
 * 描述：邢台--商家交易表
 *
 *  transaction : 交易ID( 20190122 )
 *  tra_type	: 交易类型 : 0 锘 , 1 现金
 *  status      ：0 支付, 1 退款
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS car_transaction_319;
CREATE TABLE IF NOT EXISTS car_transaction_319(
	id		INT UNSIGNED		NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '表    ID',
	tran_id		CHAR(60)		NOT NULL		COMMENT '交 易 ID',
	user_id		CHAR(32)   	    	NOT NULL 		COMMENT '用 户 ID',
	user_phone	CHAR(11)   	    	NOT NULL 		COMMENT '用户手机号',
	user_name	VARCHAR(64)   	    	NOT NULL 		COMMENT '用户姓名',
	cashier_id	CHAR(32)		NOT NULL		COMMENT '收 银 ID',
	merchant_id	CHAR(32)   	    	NOT NULL 		COMMENT '商 家 ID',
	rate		DOUBLE(10,2)		NOT NULL		COMMENT '本单费率金额',
	amount		DOUBLE(10,2)		NOT NULL		COMMENT '交易金额',
	refund_amount	DOUBLE(10,2)		NOT NULL		COMMENT '退款金额',
	mer_balance	DOUBLE(14,2)		NOT NULL		COMMENT '商户余额',
	tran_type	INT UNSIGNED		NOT NULL		COMMENT '交易类型',
	note_test	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '备    注',
	status		INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT '当前状态',
	create_at	INT UNSIGNED		NOT NULL		COMMENT	'创建时间',
	update_at	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'更新时间',
	INDEX car_transaction_319( user_id, merchant_id )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='邢台商家交易表';