
DROP TABLE IF EXISTS car_merchant_310;
/*
 * 描述：邯郸--商家信息表
 *
 *  type_id     : 商家所属行业ID
 *  status      ：0 普通商家  1 认证中 2 认证未通过 3 认证通过， 4 活动结束, 5 删除 
 *  rate	: 0 费       率
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS car_merchant_310(
	id		INT UNSIGNED		NOT NULL AUTO_INCREMENT PRIMARY KEY	COMMENT '表ID',
	fid		CHAR(32)		NOT NULL DEFAULT ''	COMMENT '总店ID',
	user_id		CHAR(32)   	    	NOT NULL 		COMMENT '用户ID',
	merchant_id	CHAR(32)   	    	NOT NULL 		COMMENT '商家ID',
	invite_code	CHAR(6)   	    	NOT NULL 		COMMENT '商家邀请码',
	merchant_type	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'商 家 类 型',
	area_number	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'市   I   D',
	area_id		INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'区   I   D',
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
	business	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '营业执照',
	number_id	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '身份证证照',
	industry	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '行业许可证',
	longitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '经       度',
	latitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '纬       度',
	cash		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '现       金',
	trust		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '鍩       分',
	credits		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '积       分',
	rate		DOUBLE(11,2)            NOT NULL DEFAULT 0.0    COMMENT '费       率',
	INDEX car_merchant_310( user_id,area_number,merchant_id,merchant_type)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='邯郸商家信息表';

DROP TABLE IF EXISTS car_merchant_319;
/*
 * 描述：邢台--商家信息表
 *
 *  type_id     : 商家所属行业ID
 *  status      ：0 普通商家  1 认证中 2 认证未通过 3 认证通过， 4 活动结束, 5 删除 
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS car_merchant_319(
	id		INT UNSIGNED		NOT NULL AUTO_INCREMENT PRIMARY KEY	COMMENT '表ID',
	fid		CHAR(32)		NOT NULL DEFAULT ''	COMMENT '总店ID',
	user_id		CHAR(32)   	    	NOT NULL 		COMMENT '用户ID',
	merchant_id	CHAR(32)   	    	NOT NULL 		COMMENT '商家ID',
	invite_code	CHAR(6)   	    	NOT NULL 		COMMENT '商家邀请码',
	merchant_type	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'商 家 类 型',
	area_number	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'市   I   D',
	area_id		INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'区   I   D',
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
	business	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '营业执照',
	number_id	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '身份证证照',
	industry	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '行业许可证',
	longitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '经       度',
	latitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '纬       度',
	cash		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '现       金',
	trust		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '鍩       分',
	credits		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '积       分',
	rate            DOUBLE(11,2)            NOT NULL DEFAULT 0.0    COMMENT '费       率',
	INDEX car_merchant_319( user_id,area_number,merchant_id,merchant_type)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='邢台商家信息表';

DROP TABLE IF EXISTS car_merchant_311;
/*
 * 描述：石家庄--商家信息表
 *
 *  type_id     : 商家所属行业ID
 *  status      ：0 普通商家  1 认证中 2 认证未通过 3 认证通过， 4 活动结束, 5 删除 
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS car_merchant_311(
	id		INT UNSIGNED		NOT NULL AUTO_INCREMENT PRIMARY KEY	COMMENT '表ID',
	fid		CHAR(32)		NOT NULL DEFAULT ''	COMMENT '总店ID',
	user_id		CHAR(32)   	    	NOT NULL 		COMMENT '用户ID',
	merchant_id	CHAR(32)   	    	NOT NULL 		COMMENT '商家ID',
	invite_code	CHAR(6)   	    	NOT NULL 		COMMENT '商家邀请码',
	merchant_type	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'商 家 类 型',
	area_number	INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'区   I   D',
	area_id		INT UNSIGNED		NOT NULL DEFAULT 0	COMMENT	'区   I   D',
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
	business	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '营业执照',
	number_id	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '身份证证照',
	industry	VARCHAR(1024)		NOT NULL DEFAULT ''	COMMENT '行业许可证',
	longitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '经       度',
	latitude	DOUBLE(17,11)		NOT NULL DEFAULT 0.0	COMMENT '纬       度',
	cash		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '现       金',
	trust		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '鍩       分',
	credits		DOUBLE(11,2)		NOT NULL DEFAULT 0.0	COMMENT '积       分',
	rate            DOUBLE(11,2)            NOT NULL DEFAULT 0.0    COMMENT '费       率',
	INDEX car_merchant_311( user_id,area_number,merchant_id,merchant_type)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='石家庄商家信息表';


