create table if not exists payments (
	id bigserial primary key,
    createdAt timestamp default current_timestamp,
    paymentNumber varchar,
    amount numeric(4,2) not null,
    paidBy bigserial not null
);

create table if not exists subscriptions (
	id bigserial primary key,
    validUntil timestamp,
	linkedAccountIDs integer[],
	subscriptionType varchar,
    accountID bigserial not null
);

alter table subscriptions add column paymentID bigserial;