drop table if exists posts cascade;
drop table if exists comment;

create table posts (
    id serial primary key,
    content text,
    author varchar(255)
);

create table comment (
    id serial primary key,
    content text,
    author varchar(255),
    post_id integer references posts(id)
);