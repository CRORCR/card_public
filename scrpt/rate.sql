

/*
 * 描述：云握费率表
 *
 ********************************************************************************************/
DROP TABLE IF EXISTS yoawo_rate;
CREATE TABLE IF NOT EXISTS yoawo_rate(
        id              INT UNSIGNED    AUTO_INCREMENT PRIMARY KEY  COMMENT '表    ID',
	class_id	INT		NOT NULL DEFAULT 0	COMMENT '一 类 ID',
	level_id	INT		NOT NULL		COMMENT '二 类 ID',
	name		CHAR(32)	NOT NULL		COMMENT '名 称 ID',
	rate_1		DOUBLE(5,2)	NOT NULL		COMMENT '费    率',
	rate_2		DOUBLE(5,2)	NOT NULL		COMMENT '费    率',
	rate_3		DOUBLE(5,2)	NOT NULL		COMMENT '费    率',
	rate_4		DOUBLE(5,2)	NOT NULL		COMMENT '费    率',
	rate_5		DOUBLE(5,2)	NOT NULL		COMMENT '费    率',
	rate_6		DOUBLE(5,2)	NOT NULL		COMMENT '费    率',
	rate_7		DOUBLE(5,2)	NOT NULL		COMMENT '费    率'
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='云握费率表';
