drop table if exists marshal_test;
create table if not exists marshal_test
(
    id            long,
    bool_value    boolean,
    float64_value decimal(4, 2),
    int64_value   long,
    string_value  varchar(32),
    time_value    timestamp
);

insert into marshal_test
values (0, true, 1.2, 3, '456', now()),
       (1, null, 1.2, 3, '456', now()),
       (2, true, null, 3, '456', now()),
       (3, true, 1.2, null, '456', now()),
       (4, true, 1.2, 3, null, now()),
       (5, true, 1.2, 3, '456', null);