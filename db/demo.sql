CREATE DATABASE IF NOT EXISTS `trans_demo`;
USE `trans_demo`;
CREATE TABLE `user`
(
    `user_id`   INT AUTO_INCREMENT COMMENT '用户ID',
    `user_name` VARCHAR(50)  NOT NULL DEFAULT '' COMMENT '用户名',
    `password`  VARCHAR(100) NOT NULL DEFAULT '' COMMENT '密码',
    `email`     VARCHAR(100) NOT NULL DEFAULT '' COMMENT '邮箱',
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '用户信息表';

CREATE TABLE `product`
(
    `product_id`     INT AUTO_INCREMENT COMMENT '商品ID',
    `product_name`   VARCHAR(100) NOT NULL DEFAULT '' COMMENT '商品名称',
    `stock_quantity` INT          NOT NULL DEFAULT 0 COMMENT '库存量',
    `version`        INT          NOT NULL DEFAULT 0 COMMENT '版本号',
    PRIMARY KEY (`product_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '商品信息表';


CREATE TABLE `order_table`
(
    `order_id`     INT AUTO_INCREMENT COMMENT '订单ID',
    `user_id`      INT NOT NULL DEFAULT 0 COMMENT '用户ID',
    `product_id`   INT NOT NULL DEFAULT 0 COMMENT '商品ID',
    `order_time`   INT NOT NULL DEFAULT 0 COMMENT '下单时间',
    `order_status` INT NOT NULL DEFAULT 0 COMMENT '订单状态',
    `order_amount` INT NOT NULL DEFAULT 0 COMMENT '下单数量',
    PRIMARY KEY (`order_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '订单信息表';


INSERT INTO `user` (user_name, `password`, email) VALUES
                                                    ('user1', 'password1', 'user1@example.com'),
                                                    ('user2', 'password2', 'user2@example.com'),
                                                    ('user3', 'password3', 'user3@example.com');
INSERT INTO `product` (product_name, stock_quantity, version) VALUES
                                                                ('Product A', 100, 1),
                                                                ('Product B', 150, 1),
                                                                ('Product C', 200, 1);
INSERT INTO `order_table` (user_id, product_id, order_time, order_status, order_amount) VALUES
                                                                                         (1, 1, UNIX_TIMESTAMP(), 1, 5),
                                                                                         (2, 2, UNIX_TIMESTAMP(), 1, 3),
                                                                                         (3, 3, UNIX_TIMESTAMP(), 1, 2);
