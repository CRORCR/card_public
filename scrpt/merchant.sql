
DROP TABLE IF EXISTS car_merchant_310;

/*
 * 描述：商家信息表
 *
 *  type_id     : 商家所属行业ID
 *  status      ：0 普通商家  1 认证中 2 认证未通过 3 认证通过， 4 活动结束, 5 删除 
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS car_merchant_310(
	id		INT UNSIGNED		NOT NULL		COMMENT '表ID',
	fid		CHAR(32)		NOT NULL DEFAULT ''	COMMENT '总店ID',
	user_id		CHAR(32)   	    	NOT NULL 		COMMENT '用户ID',
	merchant_id	CHAR(32)   	    	NOT NULL 		COMMENT '商家ID',
	merchant_type	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'商 家 类 型',
	trust_status	BOOLEAN			NOT NULL DEFAULT TRUE	COMMENT '是否支持诺',
	area_number	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'地 区 I   D',
	create_at	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'创 建 时 间',
    	describea   	VARCHAR(526)        	NOT NULL DEFAULT ''     COMMENT '文 本 描 述',
    	address   	VARCHAR(256)        	NOT NULL DEFAULT ''     COMMENT '地 址 描 述',
	name		VARCHAR(256)		NOT NULL		COMMENT '名       称',
    	status      	TINYINT UNSIGNED    	NOT NULL DEFAULT 0      COMMENT '状       态',
	phone		CHAR(11)		NOT NULL		COMMENT '手  机   号',
	icon		VARCHAR(256)		NOT NULL DEFAULT ''	COMMENT '商 家 头 像',
	loopimg    	VARCHAR(1000)           NOT NULL DEFAULT ''     COMMENT '商家 轮播图',
	infoimg    	VARCHAR(1000)           NOT NULL DEFAULT ''     COMMENT '商家 详情图',
	video    	VARCHAR(120)            NOT NULL DEFAULT ''     COMMENT '商家视频介绍',
	checkdesc    	VARCHAR(256)            NOT NULL DEFAULT ''     COMMENT '认证未失败描述',
	checkimg	VARCHAR(1280)		NOT NULL DEFAULT ''	COMMENT '认       证',
	longitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '经       度',
	latitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '纬       度',
	cash		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '现       金',
	trust		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '鍩       分',
	credits		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '积       分',
	INDEX car_merchant(id,user_id,area_number,merchant_id,merchant_type)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 COMMENT='商家信息表';

