-- csr definition

CREATE TABLE `biz_qas` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `question` varchar(355) DEFAULT NULL,
  `answer` varchar(555) DEFAULT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `biz_qa_types` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `q_type` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

insert into biz_qa_types (q_type) values ("商品问题");
insert into biz_qa_types (q_type) values ("订单问题");
insert into biz_qa_types (q_type) values ("支付问题");
insert into biz_qa_types (q_type) values ("活动问题");
insert into biz_qa_types (q_type) values ("物流问题");
insert into biz_qa_types (q_type) values ("售后问题");
insert into biz_qa_types (q_type) values ("商品使用问题");
insert into biz_qa_types (q_type) values ("店铺相关问题");
insert into biz_qa_types (q_type) values ("其它");


CREATE TABLE `biz_question_types` (
  `qid` int unsigned NOT NULL,
  `type_id` int unsigned NOT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `biz_qa_questions` (
  `qid` int unsigned NOT NULL,
  `question` varchar(355) DEFAULT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;















CREATE TABLE `biz_merchant` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `alias` varchar(255) DEFAULT NULL,
  `picture_link` varchar(255) DEFAULT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL,
   PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `biz_merchant_link` (
  `mid` int unsigned NOT NULL,
  `link` varchar(555) DEFAULT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `biz_m_tag` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `tag` varchar(155) DEFAULT NULL,
  `detail` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `biz_merchant_tag` (
  `mid` int unsigned NOT NULL,
  `tag_id` int unsigned NOT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `biz_merchant_parameters` (
  `mid` int unsigned NOT NULL,
  `parameter_name` varchar(255) DEFAULT NULL,
  `parameter_value` varchar(255) DEFAULT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `biz_merchant_keypoints` (
  `mid` int unsigned NOT NULL,
  `keypoint` varchar(355) DEFAULT NULL,
  `update_time` datetime  DEFAULT NULL,
  `updated_by` varchar(255) DEFAULT NULL 
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;









-- csr.biz_messages definition

CREATE TABLE `biz_messages` (
  `m_time` datetime NOT NULL,
  `direction` int NOT NULL,  -- 1 user question 2 assistant 
  `user_nick` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `csr_nick` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `content` varchar(555) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `url_link` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `template_id` int DEFAULT NULL,
  PRIMARY KEY (`m_time`,`user_nick`, `direction`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;


CREATE TABLE `biz_qa_hits` (
  `m_time` datetime NOT NULL,
  `user_nick` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `question` varchar(555) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`m_time`,`user_nick`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
