


/*
 * 描述：用户管理员表
 *	start	  : 起始增长率 -- 鸡的初始化增长率
 *                            
 *	limit	  : 最高增长率 -- 最高不能超过的增长率
 *
 *	status	  : 状    态   -- 0 正常状态， 1 下架
 * 
 *	lim_count : 限制销售数量   -- 0 表示不限制数量
 *
 *	lim_time  : 限制销售时间   -- 0 表示不限制时间
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS Chicken(
    id          	INT UNSIGNED        AUTO_INCREMENT PRIMARY KEY  COMMENT '鸡种类ID',
    type        	TINYINT UNSIGNED    NOT NULL                    COMMENT '鸡的类型',
    name        	CHAR(15)            NOT NULL                    COMMENT '鸡的名称',
    price               DOUBLE(9,2)         NOT NULL		        COMMENT '价    格',
    coun		INT UNSIGNED	    NOT NULL DEFAULT 0		COMMENT '已售数量',
    start               DOUBLE(9,2)         NOT NULL		        COMMENT '起始增长率',
    limit               DOUBLE(9,2)         NOT NULL		        COMMENT '最高增长率',
    create_at		INT UNSIGNED	    NOT NULL			COMMENT '创建时间',
    status              TINYINT UNSIGNED    NOT NULL DEFAULT 0          COMMENT '状    态',
    lim_count		INT UNSIGNED	    NOT NULL DEFAULT 0		COMMENT '限制数量',
    lim_time		INT UNSIGNED	    NOT NULL DEFAULT 0		COMMENT '限制时间',
    live_at		INT UNSIGNED	    NOT NULL DEFAULT 30         COMMENT '鸡的生存时间',
    growth_at		INT UNSIGNED	    NOT NULL DEFAULT 5		COMMENT '鸡的成长时间',
    chickenurl		VARCHAR(2048)	    NOT NULL    		COMMENT '鸡的形象图',
    chickenurl		VARCHAR(2048)	    NOT NULL  			COMMENT '鸡的-蛋图',
    INDEX users( id, name )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=30000 COMMENT='鸡的种类表';


