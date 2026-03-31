-- Create tweets table

Create table tweets (
	id serial primary key,
    user_id int not null,
    tweet varchar(280) not null,
    createdAt timestamp default current_timestamp,
    updatedAt timestamp default current_timestamp
);