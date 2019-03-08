
DROP TABLE IF EXISTS car_withdrawal_account;
/*
 * 描述：商家提现账号管理表
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS car_withdrawal_account(
        id              INT UNSIGNED    AUTO_INCREMENT PRIMARY KEY  COMMENT '表    ID',
        merchant_id     CHAR(32)        NOT NULL                COMMENT '商 户 ID',
	account_type	CHAR(32)	NOT NULL		COMMENT '账号类型',
	account		CHAR(64)	NOT NULL		COMMENT '账    号',
	phone		CHAR(11)	NOT NULL		COMMENT '手 机 号',
	user_name	VARCHAR(126)	NOT NULL		COMMENT '真是姓名',
	car_type	CHAR(32)	NOT NULL		COMMENT '证件类型',
	car_number	CHAR(64)	NOT NULL		COMMENT '证件号码',
	create_at	INT		NOT NULL		COMMENT '插入时间',
	state   	INT		NOT NULL DEFAULT 0	COMMENT '当前状态',
	is_main		CHAR(5)		NOT NULL		COMMENT '是否是主提现', -- "TRUE", "FALSE"
        INDEX car_withdrawal_account( merchant_id )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='商家提现记录';

