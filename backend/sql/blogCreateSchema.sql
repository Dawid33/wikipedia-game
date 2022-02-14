create table blog.comments
(
    uuid          uuid      default uuid_generate_v4()
        constraint table_name_pk
            primary key,
    name          text not null,
    comment       text not null,
    post_time     timestamp not null default current_timestamp,
    nesting_level int not null default 0,
    parent        uuid default null,
    blog_post_id  int not null
);