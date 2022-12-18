create table task_status (
    id integer primary key,
    status varchar(20)
);

create table email_type (
    id integer primary key,
    type varchar(20)
);

create table task (
    id integer primary key,
    creation_date timestamp not null,
    author varchar(50) not null,
    status_id integer references task_status on delete cascade
);

create table link_click (
    id integer primary key,
    action_date timestamp not null,
    task_id integer references task on delete cascade,
    action_author varchar(255) not null,
    action_result boolean not null
);

create table email (
    id integer primary key,
    sent_date timestamp not null,
    address varchar(50) not null,
    task_id integer references task on delete cascade,
    email_type_id integer references email_type on delete cascade
);

