package golang_demo

// User 拥有并属于多种 language，`user_languages` 是连接表
// type Username struct {
// 	Model
// 	Name   string
// 	Emails []Email `gorm:"many2many:username_email;"`
// }

// type Email struct {
// 	Model
// 	Addr string
// }

// func TestAssc() {
// 	db := config.Con()
// 	var username Username
// 	db.Scm.Where("id = ?", 1).Find(&username)

// 	var email []Email
// 	db.Scm.Find(&email)
// 	db.Scm.Model(&username).Association("Emails").Replace(&email)
// }

// CREATE TABLE `email` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT,
// 	`addr` varchar(255) DEFAULT NULL,
// 	`created_at` int(11) DEFAULT NULL COMMENT '创建时间',
// 	`updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
// 	`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;

// CREATE TABLE `username` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT,
// 	`name` varchar(255) DEFAULT NULL,
// 	`updated_at` int(11) DEFAULT NULL COMMENT '更新时间',
// 	`deleted_at` datetime DEFAULT NULL COMMENT '删除时间',
// 	`created_at` int(11) DEFAULT NULL,
// 	PRIMARY KEY (`id`)
//   ) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

// CREATE TABLE `username_email` (
// 	`id` int(11) NOT NULL AUTO_INCREMENT,
// 	`username_id` int(11) DEFAULT NULL,
// 	`email_id` int(11) DEFAULT NULL,
// 	PRIMARY KEY (`id`),
// 	UNIQUE KEY `id` (`username_id`,`email_id`) USING BTREE
//   ) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
