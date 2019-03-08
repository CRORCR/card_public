DROP TABLE IF EXISTS car_withdrawal_foot;
/*
 * 描述：提现记录表
 *
 *      wit_type   : 提现类型 1 现金提现， 2 微信提现， 3 支付宝提现
 *      state      : 状    态 0 提交，1 审核， 2 审核通过，3 审核未通过， 6 到帐
 *
 ********************************************************************************************/
CREATE TABLE IF NOT EXISTS car_withdrawal_foot(
        id              INT UNSIGNED    AUTO_INCREMENT PRIMARY KEY  COMMENT '表    ID',
        wit_id          CHAR(48)        NOT NULL                COMMENT '提现单号',
        wit_type        INT             NOT NULL                COMMENT '提现类型',
        amount          DOUBLE(9,2)     NOT NULL                COMMENT '提现金额',
        poundage        DOUBLE(9,2)     NOT NULL                COMMENT '手 续 费',
        submission_at   INT             NOT NULL                COMMENT '提交时间',
        state           INT             NOT NULL                COMMENT '状    态',
        arrive_at       INT             NOT NULL DEFAULT 0      COMMENT '到帐时间',
        desc_info       VARCHAR(1024)   NOT NULL DEFAULT ''     COMMENT '描    述',
        user_id         CHAR(32)        NOT NULL                COMMENT '用 户 ID',
        merchant_id     CHAR(32)        NOT NULL                COMMENT '商 户 ID',
	merchant_name   VARCHAR(64)     NOT NULL                COMMENT '商户名称',
        user_name       VARCHAR(126)    NOT NULL                COMMENT '用户名称',
        user_phone      CHAR(11)        NOT NULL                COMMENT '手 机 号',
        target_id       VARCHAR(64)     NOT NULL                COMMENT '目标账户',
        INDEX car_withdrawal_foot( wit_id, user_id, amount, merchant_id, arrive_at )
)ENGINE=InnoDB DEFAULT CHARSET=UTF8 AUTO_INCREMENT=0 COMMENT='商家提现记录';

