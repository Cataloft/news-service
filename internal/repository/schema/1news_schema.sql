create table if not exists news
(
    id      serial
    constraint news_pk
    primary key,
    title   varchar not null,
    content varchar
);

alter table news
    owner to root;

create table if not exists news_categories
(
    news_id     integer not null,
    category_id integer not null,
    constraint news_categories_pk
    primary key (news_id, category_id),
    constraint news_categories_pk_unique
    unique (news_id, category_id)
    );

alter table news_categories
    owner to root;