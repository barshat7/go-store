
-- Create Category Table
create table category (
id int not null auto_increment,
name varchar(255) not null,
description varchar(255),
primary key (id)
);


-- Create Product Table
create table product (
id int not null auto_increment,
name varchar(255) not null,
description varchar(255),
available boolean,
added_date datetime,
unit_price double (10, 2),
category_id int,
primary key (id),
foreign key (category_id) references category(id)
)