

create table cursors (
   id varchar(255) not null,
   `cursor` bigint not null,
   updated_at datetime(3) not null,

   primary key (id)
);

create table events (
  id bigint not null auto_increment,
  foreign_id bigint not null,
  timestamp datetime(3) not null,
  type int not null,

  primary key (id)
);

create table rounds (
  id bigint not null auto_increment,        -- 1
  external_id bigint not null,                   -- 1
  player varchar(255),                      -- "Skunk#1"
  rank int not null,                        -- 1
  status int not null,                      -- StatusCollected(1)

  created_at datetime(3) not null,
  updated_at datetime(3) not null,

  unique by_team_status (player, external_id),

  primary key (id)
);

create table parts (
    id bigint not null auto_increment,
    round_id int not null,                 
    player varchar(255) not null,
    rank int not null,
    part int not null,

    created_at datetime not null,

    primary key(id)
);
