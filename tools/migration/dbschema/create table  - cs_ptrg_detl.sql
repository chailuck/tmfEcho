    CREATE TABLE cs_ptrg_detl (
refn_seqn integer NOT NULL,
titl varchar(40) ,
frst_name varchar(80) NOT NULL,
last_name varchar(80) ,
id_type char(2) NOT NULL,
id_numb char(20) NOT NULL,
gndr char(1) ,
date_of_brth date ,
emal_addr varchar(40) ,
ntnt_code char(4) ,
crtd_dttm TIMESTAMP(0) NOT NULL,
crtd_by char(12) NOT NULL,
last_chng_dttm TIMESTAMP(0) NOT NULL,
last_chng_by char(12) NOT NULL
);
