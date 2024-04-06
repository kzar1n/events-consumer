create table if not exists authorization.payments (
    id_payment                  varchar(36) primary key,
    id_account                  varchar(36) not null,
    payment_type                varchar(10),
    payment_date                date not null,	
    payment_value               decimal(15,2) not null
);

