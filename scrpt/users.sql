
DROP TABLE IF EXISTS Users;

/*
 * 描述：用户管理员表
 *
 *  type_id     : 用户创建类型 
 *			0 微信 
 *			1 手机号 
 *
 *  sex		: 0 未知 1 女 2 男 
 *
 *  role        ：角色：
 *                      0 普通用户
 *			1 特殊提款用户
 *			3 商户用户
 *
 *  status      ：状态：0 正常 1 禁用 2 违法操作
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS Users(
    id          	INT UNSIGNED        AUTO_INCREMENT PRIMARY KEY  COMMENT '用 户 ID',
    share_id 		CHAR(32)            NOT NULL                    COMMENT '用户分享ID',
    iconurl		VARCHAR(512)	    NOT NULL DEFAULT ''		COMMENT '头像地址',
    type_id     	TINYINT UNSIGNED    NOT NULL DEFAULT 0          COMMENT '注册类型',
    phone       	CHAR(11)            NOT NULL                    COMMENT '账    号',
    open_id     	VARCHAR(516)        NOT NULL DEFAULT ''         COMMENT '微信OpenId',
    name        	VARCHAR(126)        NOT NULL                    COMMENT '姓    名',
    token       	CHAR(32)            NOT NULL DEFAULT ''         COMMENT 'TOKEN   ',
    unionid_android  	CHAR(32)       	    NOT NULL DEFAULT ''         COMMENT 'and unionid ',
    unionid_ios  	CHAR(32)            NOT NULL DEFAULT ''         COMMENT 'ios unionid ',
    number_id   	CHAR(27)            NOT NULL DEFAULT ''         COMMENT '身份证号',
    loginpass   	CHAR(32)            NOT NULL DEFAULT ''         COMMENT '登陆密码',
    paypass     	CHAR(32)            NOT NULL DEFAULT ''         COMMENT '支付密码',
    create_at   	INT UNSIGNED        NOT NULL                    COMMENT '创建时间',
    update_at   	INT UNSIGNED        NOT NULL DEFAULT 0          COMMENT '更新时间',
    email       	VARCHAR(64)         NOT NULL DEFAULT ''         COMMENT '邮箱地址',
    role        	TINYINT UNSIGNED    NOT NULL DEFAULT 0          COMMENT '角    色',
    sex         	TINYINT UNSIGNED    NOT NULL DEFAULT 0          COMMENT '性    别',
    attesta     	TINYINT UNSIGNED    NOT NULL DEFAULT 0          COMMENT '身份认证',
    status      	TINYINT UNSIGNED    NOT NULL DEFAULT 0          COMMENT '状    态',
    cash                DOUBLE(9,2)         NOT NULL DEFAULT 0.0        COMMENT '现    金',
    trust               DOUBLE(9,2)         NOT NULL DEFAULT 0.0        COMMENT '鍩    分',
    code		CHAR(6)		    NOT NULL			COMMENT '邀 请 码',
    INDEX users(share_id,jwt_token,token,phone,unionid_android)
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=30000 COMMENT='用户信息表';


