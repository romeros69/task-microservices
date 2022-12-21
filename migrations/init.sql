create table task_status (
    id serial primary key,
    status varchar(20)
);

create table email_type (
    id serial primary key,
    type varchar(20)
);

create table task (
    id serial primary key,
    creation_date timestamp not null,
    author varchar(50) not null,
    status_id integer references task_status on delete cascade
);

create table link_click (
    id serial primary key,
    action_date timestamp not null,
    task_id integer references task on delete cascade,
    action_author varchar(255) not null,
    action_result boolean not null
);

create table email (
    id serial primary key,
    sent_date timestamp not null,
    address varchar(50) not null,
    task_id integer references task on delete cascade,
    email_type_id integer references email_type on delete cascade
);

-- create or replace function check_task()
-- returns trigger as $$
-- begin
--
-- end;
-- $$ language 'plpgsql';



insert into task_status (id, status) values (1, 'in_progress');
insert into task_status (id, status) values (2, 'finished_rejected');
insert into task_status (id, status) values (3, 'finished_approved');

insert into email_type (id, type) values (1, 'agreement');
insert into email_type (id, type) values (2, 'end_negative');
insert into email_type (id, type) values (3, 'agreement_update');
