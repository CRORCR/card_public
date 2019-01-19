
/*
 * 描述：商家员工表字段说明
 * =======================================================================================
 * authority: 
 *	0:1 收银权限
 *	如果员工存在收银权限: 收款码的生成规则
 * 		{
 *			"M": "12345678901234567890123456789012",
 *			"U": "12345678901234567890123456789012"
 *		}
 *	M: 商家ID
 *	U: 用户ID
 *-----------------------------------------------------------------------------------------
 * sex: true 男, false 女
 *-----------------------------------------------------------------------------------------
 * number_fage: 1: 店长
 *
 ********************************************************************************************/

/*
 * 描述：河北-邯郸商家员工表
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS chi_staff_310;
CREATE TABLE IF NOT EXISTS chi_staff_310(
        id              INT UNSIGNED    AUTO_INCREMENT PRIMARY KEY  COMMENT '表    ID',
	merchant_id	CHAR(32)	NOT NULL		COMMENT '商 家 ID',
	user_id		CHAR(32)	NOT NULL DEFAULT ''	COMMENT '员 工 ID',
        name	   	CHAR(50)        NOT NULL                COMMENT '姓    名',
        phone	   	CHAR(11)        NOT NULL                COMMENT '手 机 号',
	number_id	CHAR(18)	NOT NULL DEFAULT ''	COMMENT '身份证号',
	sex		BOOLEAN		NOT NULL 		COMMENT '性    别',
	create_at	INT UNSIGNED    NOT NULL		COMMENT '创建时间',
	state		INT UNSIGNED	NOT NULL		COMMENT '状    态',
	number_fage	INT UNSIGNED 	NOT NULL		COMMENT '身份标识',
	authority       INT UNSIGNED	NOT NULL DEFAULT 0	COMMENT '权    限',
        INDEX chi_staff_310( merchant_id, user_id, phone, number_fage )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='河北-邯郸商家员工表';

/*
 * 描述：河北-邢台商家员工表
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS chi_staff_319;
CREATE TABLE IF NOT EXISTS chi_staff_319(
        id              INT UNSIGNED    AUTO_INCREMENT PRIMARY KEY  COMMENT '表    ID',
	merchant_id	CHAR(32)	NOT NULL		COMMENT '商 家 ID',
	user_id		CHAR(32)	NOT NULL DEFAULT ''	COMMENT '员 工 ID',
        name	   	CHAR(50)        NOT NULL                COMMENT '姓    名',
        phone	   	CHAR(11)        NOT NULL                COMMENT '手 机 号',
	number_id	CHAR(18)	NOT NULL DEFAULT ''	COMMENT '身份证号',
	sex		BOOLEAN		NOT NULL 		COMMENT '性    别',
	create_at	INT UNSIGNED    NOT NULL		COMMENT '创建时间',
	state		INT UNSIGNED	NOT NULL		COMMENT '状    态',
	number_fage	INT UNSIGNED 	NOT NULL		COMMENT '身份标识',
	authority       INT UNSIGNED	NOT NULL DEFAULT 0	COMMENT '权    限',
        INDEX chi_staff_319( merchant_id, user_id, phone, number_fage )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='河北-邢台商家员工表';

/*
 * 描述：河北-石家庄商家员工表
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS chi_staff_311;
CREATE TABLE IF NOT EXISTS chi_staff_311(
        id              INT UNSIGNED    AUTO_INCREMENT PRIMARY KEY  COMMENT '表    ID',
	merchant_id	CHAR(32)	NOT NULL		COMMENT '商 家 ID',
	user_id		CHAR(32)	NOT NULL DEFAULT ''	COMMENT '员 工 ID',
        name	   	CHAR(50)        NOT NULL                COMMENT '姓    名',
        phone	   	CHAR(11)        NOT NULL                COMMENT '手 机 号',
	number_id	CHAR(18)	NOT NULL DEFAULT ''	COMMENT '身份证号',
	sex		BOOLEAN		NOT NULL 		COMMENT '性    别',
	create_at	INT UNSIGNED    NOT NULL		COMMENT '创建时间',
	state		INT UNSIGNED	NOT NULL		COMMENT '状    态',
	number_fage	INT UNSIGNED 	NOT NULL		COMMENT '身份标识',
	authority       INT UNSIGNED	NOT NULL DEFAULT 0	COMMENT '权    限',
        INDEX chi_staff_311( merchant_id, user_id, phone, number_fage )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='河北-石家庄商家员工表';

