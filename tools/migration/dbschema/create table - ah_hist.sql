create table AH_HIST (
    cust_numb integer NOT NULL,
    subr_numb char(12) NOT NULL,
    acty_code char(8) NOT NULL,
    crtd_dttm TIMESTAMP(0) NOT NULL,
    crtd_by char(12) NOT NULL,
    hist_desc1 char(30) NOT NULL,
    hist_desc2 char(30),
    anls_code1 char(12),
    anls_code2 char(12),
    anls_code3 char(12)
)