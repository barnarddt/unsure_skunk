create table skunk_cursors (
   id varchar(255) not null,
   `cursor` bigint not null,
   updated_at datetime(3) not null,

   primary key (id)
);

create table skunk_events (
  id bigint not null auto_increment,
  foreign_id bigint not null,
  timestamp datetime(3) not null,
  type int not null,

  primary key (id)
);

create table skunk_rounds (
  id bigint not null auto_increment,
  round_id bigint not null,
  player varchar(255),
  status int not null,

  created_at datetime(3) not null,
  updated_at datetime(3) not null,

  primary key (id),
  unique by_team_status (player,round_id)
);
