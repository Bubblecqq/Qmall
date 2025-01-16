CREATE DATABASE mall_seckill;
USE mall_seckill;

create table `mall_seckill_stock_t` (
                                   `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                                   `products_id`                  bigint(20)   comment '商品ID',
                                   `stock`                 int(11)      not null                comment '库存大小',
                                   `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                   `update_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `idx_productsid` (`products_id`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀库存表' ;



create table `mall_seckill_record_t` (
                               `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                               `user_id`                  bigint(20)  not null comment '用户ID',
                               `products_id`                  bigint(20) not null  comment '商品ID',
                               `sec_num`               varchar(128)                         comment '秒杀号',
                               `order_num`               varchar(128)                         comment '订单号',
                               `price`                 int(11)      not null                comment '金额',
                               `status`                 int(11)      not null                comment '状态',
                               `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                               `update_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                               PRIMARY KEY (`id`),
                               UNIQUE KEY `idx_secnum` (`sec_num`),
                               UNIQUE KEY `idx_ordernum` (`order_num`),
                               KEY `idx_userid` (`user_id`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀记录表' ;


create table `mall_products_t` (
                                    `id`                  bigint(20)       not null AUTO_INCREMENT comment '',
                                    `products_num`               varchar(128)                         comment '商品编号',
                                    `products_name`               varchar(128)                         comment '商品名字',
                                    `price`                 float      not null                comment '价格',
                                    `picture_url`               varchar(128)                         comment '商品图片',
                                    `seller`                  bigint(20)  not null comment '卖家ID',
                                    `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                    `update_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                    PRIMARY KEY (`id`),
                                    UNIQUE KEY `idx_productsnum` (`products_num`),
                                    KEY `idx_seller` (`seller`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀商品表' ;

create table `mall_order_t` (
                                    `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                                    `seller`                  bigint(20)  not null comment '买方ID',
                                    `buyer`                  bigint(20)  not null comment '卖方ID',
                                    `products_id`                  bigint(20) not null  comment '商品ID',
                                    `products_num`               varchar(128)                         comment '商品编号',
                                    `order_num`               varchar(128)                         comment '订单号',
                                    `price`                 int(11)      not null                comment '金额',
                                    `status`                 int(11)      not null                comment '状态',
                                    `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                    `update_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                    PRIMARY KEY (`id`),
                                    KEY `idx_productsid` (`products_id`),
                                    UNIQUE KEY `idx_ordernum` (`order_num`),
                                    KEY `idx_seller` (`seller`),
                                    KEY `idx_buyer` (`buyer`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀订单表' ;


insert into mall_products_t(products_name, products_num, price, picture_url, seller) values("redhat", "abc123", 18, "http://", 135);

insert into mall_seckill_stock_t(products_id, stock) values(1, 3);

create table `mall_quota_t` (
                                   `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                                   `products_id`                  bigint(20)   comment '商品ID',
                                   `num`                 int(11)      not null                comment '限额',
                                   `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                                   `update_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                                   PRIMARY KEY (`id`),
                                   UNIQUE KEY `idx_productsid` (`products_id`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀限额表' ;

create table `mall_user_quota_t` (
                           `id`                  bigint(20)       not null AUTO_INCREMENT comment 'ID',
                           `user_id`                  bigint(20)  not null comment '用户ID',
                           `products_id`                  bigint(20)   comment '商品ID',
                           `num`                 int(11)      not null         comment '限额',
                           `killed_num`          int(11)      not null         comment '已经消耗的额度',
                           `create_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP                             comment '创建时间',
                           `update_time`         datetime     not null DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP comment '修改时间',
                           PRIMARY KEY (`id`),
                           KEY `idx_productsid` (`products_id`),
                           KEY `idx_userproductsid` (`user_id`, `products_id`)
)ENGINE=InnoDB  default CHARSET=utf8mb4 comment '秒杀用户限额表' ;