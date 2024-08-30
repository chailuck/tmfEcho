create table rd_slmn_detl (
deal_numb            char(8) NOT NULL,
slmn_code            char(5) NOT NULL,
effc_date            date NOT NULL,
expr_date            date,
crtd_dttm            timestamp(0) NOT NULL,
crtd_by              char(12) NOT NULL,
last_chng_dttm       timestamp(0) NOT NULL,
last_chng_by         char(12) NOT NULL
)
