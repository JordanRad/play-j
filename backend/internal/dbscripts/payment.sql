create table if not exists payment_details (
	id bigserial primary key,
    cvv varchar,
    bic varchar,
    cardNumber varchar,
    expirationDate varchar,
    accountID bigserial not null
);

create table if not exists payments (
	id bigserial primary key,
    createdAt timestamp default current_timestamp,
    paymentNumber varchar,
    paidBy bigserial not null
);

create table if not exists subscriptions (
	id bigserial primary key,
    validUntil timestamp,
	linkedAccountIDs integer[],
	subscriptionType varchar,
    accountID bigserial not null
);