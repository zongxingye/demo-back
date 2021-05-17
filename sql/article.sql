/* 创建文章表*/
create table articles ( `id` int unsigned auto_increment,
 `user_id` varchar(20) not null,
 `contents` varchar(255) default null,
 `title` varchar(20) default null,
 `spot` int default null,
 `create_at` timestamp not null default current_timestamp comment ' time of the message created',
 `update_at` datetime default null on update current_timestamp comment 'modify time',
  primary key(`id`)
)
engine = InnoDB auto_increment=1 default charset = utf8