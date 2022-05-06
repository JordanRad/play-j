insert into artists (id,"name",storagefoldername) values 
(1,'Metallica','Metallica'),
(2,'Drake','Drake'),
(3,'Travis Scott','TravisScott');

insert into tracks (id,"name",fullname,storagelocation,createdat) values 
(1,'the ends','The ends - Travis Scott ','the_ends.mp3','2022-05-01 12:54:30.553'),
(2,'pick up the phone','Pick up the phone - Travis Scott ','pick_up_the_phone.mp3','2022-05-01 12:54:30.553'),
(3,'Im upset','Im upset - Drake','im_upset.mp3','2022-05-01 12:54:30.553'),
(4,'One','Metallica - One','Mettalica/One.mp3','2022-04-22 18:03:12.972'),
(5,'nonstop','Nonstop - Drake ','Drake - Nonstop.mp3','2022-05-01 12:54:30.553');

insert into albums (id,"name",genre,trackids,createdat) values 
(1,'And Justice for all','Metal','{4}','2022-04-22 18:02:25.792'),
(2,'Scorpion','Rap','{3,5}','2022-05-01 12:53:13.446'),
(3,'Birds in the trap','Trap','{1,2}','2022-05-01 12:53:13.446');