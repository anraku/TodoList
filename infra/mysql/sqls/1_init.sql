CREATE DATABASE app;
use app;

CREATE TABLE tasks (
  id int(11) unsigned not null auto_increment,
  text varchar(255) not null,
  done tinyint(1) not null DEFAULT 0,
  created_at datetime not null default current_timestamp,
  updated_at datetime not null default current_timestamp on update current_timestamp,
  primary key (id)
);
