CREATE TABLE cs_sms_barr (
cust_numb integer NOT NULL,
subr_numb char(12) NOT NULL,
blpd_code char(2) NOT NULL,
barr_flag char(1) NOT NULL,
crtd_dttm TIMESTAMP(0) NOT NULL,
crtd_by char(12) NOT NULL,
last_chng_dttm TIMESTAMP(0) NOT NULL,
last_chng_by char(12) NOT NULL
)
