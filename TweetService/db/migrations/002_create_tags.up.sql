-- Create tags table

Create table tags(
	id serial primary key,
    name varchar(255) unique not null
);

-- Create association table between tweets and tags
Create table tweet_tags(
	tweet_id  bigint unsigned not null,
    tag_id bigint unsigned not null,
    primary key(tweet_id,tag_id),
    foreign key (tweet_id) references tweets(id),
    foreign key (tag_id) references tags(id)
);