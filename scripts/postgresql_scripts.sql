CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

create table if not exists payments (
    id_payment                  varchar(36) primary key,
    id_account                  varchar(36) not null,
    payment_type                varchar(10),
    payment_date                date not null,    
    payment_value               decimal(15,2) not null,
    payment_status              varchar(10) default 'PENDING',
    created_at                  timestamp not null default current_timestamp,
    updated_at                  timestamp not null default current_timestamp,
    update_count                integer not null default 0
);

INSERT INTO payments (id_payment, id_account, payment_type, payment_date, payment_value, payment_status) VALUES
(uuid_generate_v4(), uuid_generate_v4(), 'CREDIT', '2022-01-01', 100.00, 'PENDING'),
(uuid_generate_v4(), uuid_generate_v4(), 'DEBIT', '2022-01-02', 200.00, 'COMPLETED'),
(uuid_generate_v4(), uuid_generate_v4(), 'CREDIT', '2022-01-03', 300.00, 'FAILED'),
(uuid_generate_v4(), uuid_generate_v4(), 'DEBIT', '2022-01-04', 400.00, 'PENDING'),
(uuid_generate_v4(), uuid_generate_v4(), 'CREDIT', '2022-01-05', 500.00, 'COMPLETED'),
(uuid_generate_v4(), uuid_generate_v4(), 'DEBIT', '2022-01-06', 600.00, 'FAILED'),
(uuid_generate_v4(), uuid_generate_v4(), 'CREDIT', '2022-01-07', 700.00, 'PENDING'),
(uuid_generate_v4(), uuid_generate_v4(), 'DEBIT', '2022-01-08', 800.00, 'COMPLETED'),
(uuid_generate_v4(), uuid_generate_v4(), 'CREDIT', '2022-01-09', 900.00, 'FAILED'),
(uuid_generate_v4(), uuid_generate_v4(), 'DEBIT', '2022-01-10', 1000.00, 'PENDING'),
(uuid_generate_v4(), uuid_generate_v4(), 'CREDIT', '2022-01-11', 1100.00, 'COMPLETED'),
(uuid_generate_v4(), uuid_generate_v4(), 'DEBIT', '2022-01-12', 1200.00, 'FAILED'),
(uuid_generate_v4(), uuid_generate_v4(), 'CREDIT', '2022-01-13', 1300.00, 'PENDING'),
(uuid_generate_v4(), uuid_generate_v4(), 'DEBIT', '2022-01-14', 1400.00, 'COMPLETED');