CREATE DATABASE  IF NOT EXISTS `yugong` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `yugong`;

DROP TABLE IF EXISTS `report`;
CREATE TABLE `report` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT '',
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `create_user` varchar(255) NOT NULL DEFAULT '',
  `network_id` bigint(20) NOT NULL,
  `status` enum('INIT','ACTIVE','INACTIVE') NOT NULL DEFAULT 'INIT',
  `last_run_at` timestamp NULL DEFAULT NULL,
  `date_range` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `uniq_report_name` (`network_id`,`create_user`,`name`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;

DROP TABLE IF EXISTS `report_message`;
CREATE TABLE `report_message` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `report_id` bigint(20) DEFAULT NULL,
  `run_mode` enum('MANUAL','SCHEDULE','SCHEDULE_BIG','SCHEDULE_DEDICATE','ANONYMOUS','ONDEMAND','ONDEMAND_BIG','DEFAULT') NOT NULL,
  `status` enum('INIT','LOAD','ACTIVE','INACTIVE','CANCEL') NOT NULL,
  `start_time` timestamp NULL DEFAULT NULL,
  `end_time` timestamp NULL DEFAULT NULL,
  `task_id` varchar(255) DEFAULT NULL,
  `report_file_option` varchar(60) DEFAULT '' COMMENT 'file,csv,xls,pdf\\nComma separated',
  `file_path_base` varchar(512) DEFAULT '',
  `local_storage_path` varchar(512) DEFAULT NULL,
  `row_count` bigint(20) NOT NULL DEFAULT '0',
  `file_size` bigint(20) NOT NULL DEFAULT '0',
  `filters` varchar(4096) DEFAULT NULL,
  `date_range` varchar(255) DEFAULT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  `retry_count` tinyint(4) DEFAULT '0',
  `column_header` text,
  `network_id` bigint(20) NOT NULL,
  `create_user` varchar(255) NOT NULL DEFAULT '',
  `error_detail` varchar(4096) DEFAULT NULL,
  `error_summary` varchar(255) DEFAULT NULL,
  `ip` varchar(128) DEFAULT NULL,
  `network_function` varchar(2048) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  KEY `idx_report_message_rptid` (`report_id`),
  KEY `idx_msg_nid_cuser` (`network_id`,`create_user`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 ROW_FORMAT=DYNAMIC;