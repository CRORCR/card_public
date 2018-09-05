


/*
 * 描述：用户管理员表
 *
 *	ntype	: 0 省， 1 直辖市， 2 自治区， 3 特别行政区
 *      open	: 0 未开放， 开放后变为时间戳
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS China(
 	ntype	     INT UNSIGNED	NOT NULL DEFAULT 0		COMMENT '省市类别',
 	open_at      INT UNSIGNED	NOT NULL DEFAULT 0		COMMENT '开放时间',
	addname	     VARCHAR(255)	NOT NULL			COMMENT	'地方名称',
	addnumb	     CHAR(15)		NOT NULL			COMMENT	'地方编号',
    	INDEX users(addnumb,addnume)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 COMMENT='中国地理编码表';


