insert into payments (id,createdat,paymentnumber,amount, paidBy) values 
(1,'2022-04-25 13:16:56.509','1650892613-PJ1-ef711045-0a6b-49f7-8ef1-8e4973b72865',11.99,1),
(2,'2022-04-26 13:31:42.959','1650893499-PJ1-de4a4ed0-7135-4317-82e9-42fdbf35867d',11.99,1),
(3,'2022-05-25 20:38:43.448','1651091923-PJ1-3b669d5b-4dfd-4749-9d15-299b323edbd1',11.99,2),
(4,'2022-05-26 20:18:50.895','1651091930-PJ1-dc2349ee-f836-4626-acf2-cbf55d574e4c',11.99,2);

insert into subscriptions (id,validuntil,linkedaccountids,subscriptiontype,accountID) values 
(1,'2022-03-25 10:40:00.000','{10,20,30}','family',1),
(2,'2022-04-25 09:30:00.000','{10,20,30}','family',2);
