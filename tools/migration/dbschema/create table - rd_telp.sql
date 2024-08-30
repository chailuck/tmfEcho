drop table "C##OMDTAC".rd_telp;

create table "C##OMDTAC".rd_telp (
telp_numb VARCHAR(12) NOT NULL,
blpd_code VARCHAR(2) NOT NULL,
last_4dgt INT NOT NULL,
old_imsi_grup INT ,
imsi_grup INT ,
area_code VARCHAR(1) ,
allc_type VARCHAR(1) NOT NULL,
brnc_code VARCHAR(2) NOT NULL,
deal_numb VARCHAR(8) ,
gold_type VARCHAR(8) ,
spcl_type VARCHAR(5) ,
live_flag VARCHAR(1) NOT NULL,
resv_flag VARCHAR(1) ,
card_numb VARCHAR(19) ,
telp_type VARCHAR(3) ,
swof_date DATE ,
load_indc VARCHAR(1) NOT NULL,
used_flag VARCHAR(1) NOT NULL,
allw_ruse VARCHAR(1) NOT NULL,
dtac_flag VARCHAR(1) NOT NULL,
comp_code VARCHAR(2) NOT NULL,
netw_type VARCHAR(10) NOT NULL,
allw_revk VARCHAR(1) ,
asgn_resn_code VARCHAR(2) ,
retn_flag VARCHAR(1) ,
crtd_dttm TIMESTAMP(0) NOT NULL,
crtd_by VARCHAR(12) NOT NULL,
last_chng_dttm TIMESTAMP(0) NOT NULL,
last_chng_by VARCHAR(12) NOT NULL

);