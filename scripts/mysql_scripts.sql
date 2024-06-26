create database if not exists `authorization`;

create table if not exists authorization.payments (
    id_payment                  varchar(36) primary key,
    id_account                  varchar(36) not null,
    payment_type                varchar(10),
    payment_date                date not null,	
    payment_value               decimal(15,2) not null,
    created_at                  timestamp not null default current_timestamp,
    updated_at                  timestamp not null default current_timestamp,
    update_count                integer not null default 0
);

insert into authorization.payments (id_payment, id_account, payment_type, payment_date, payment_value) 
values (uuid(), uuid(), 'PIX', '2021-01-01', 100.00);