/* 创建user表*/
create table users(
`id` int unsigned auto_increment,
 `username` varchar(20) not null,
 `password` varchar(20) not null,
`nickname` varchar(20) not null,
 `create_at` timestamp not null default current_timestamp comment ' time of the message created',
`update_at` datetime default null on update current_timestamp comment 'modify time',
primary key (`id`)
    ) engine=InnoDB auto_increment = 1 default charset = utf8mb4 comment='blogpractice';
