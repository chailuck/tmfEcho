--mnload to cs_subr_pcn_06.dat

--select first 1000 * from cs_subr_pcn_06
--where subr_stts in ('A','S')
--into temp tmp_subr_06
--with no log;
/*
unload to cs_cust.dat
select *
from cs_cust a, tmp_subr_06 b where a.cust_numb = b.cust_numb
;


--unload to cs_pmpt_regn.dat
select *
from cs_pmpt_regn a where exists( select 1 from  tmp_subr_06 b where a.cust_numb = b.cust_numb and a.subr_numb = b.subr_numb)
into temp tmp_regn
with no log;
*/
unload to cs_ptrg_detl.dat
select *
from cs_ptrg_detl a where exists( select 1 from  tmp_regn b where a.refn_seqn = b.seqn)
;





ca_row_id       CA_NOT_FOUND
cust_numb       385003898
subr_numb       66816121360
blpd_code       02
subr_type
neo_subr_type
telp_type       T
titl            Mr.
frst_name       dtac
last_name       trinet
subr_stts       C
from_optr_code
mnpi_stts
mnpo_stts
frst_pkgp_code  11000102
pkgp_strt_date  1995-01-16
new_ppty_code   100
dpst_txtp_code  0901
conx_txtp_code  0601
telp_allc_type  A
hrdw_numb       495003100119320
card_numb       8966189410100327616
sms_lang        T
ivr_lang        T
ussd_lang       E
cc_lang         01
deal_numb       10269002
slmn_code
agmt_chck_flag  I
agmt_chck_date  2017-08-02
agmt_chck_by    LLNIRAMOSU
frst_swon_dttm
brth_swon_dttm  1995-01-16 11:16:00
swon_dttm       1995-01-16 11:16:00
swon_by         T0624
swon_resn_code  11100001
swon_area_code  1
swof_dttm       1995-05-03 10:26:29
swof_by         T0624
swof_resn_code  9000
rcnx_resn_code
leas_code
leas_strt_date
leas_expr_date
leas_acex_date
pswd_flag       1
pswd
memb_id         YYYY
remk
id_type
id_numb
gndr
occp_code
marl_stts
salr_levl
educ_levl
emal_addr
vrfy_emal_dttm
wait_emal_addr
wait_emal_dttm
rcnx_paid_flag  0
stop_bill_flag  0
rfnd_dttm
old_cust_numb
old_subr_numb
new_cust_numb
new_subr_numb
go_inter_flag   N
chwr_cont_aou   0
last_chng_dttm  2017-08-02 11:58:03
last_chng_by    LLNIRAMOSU




Column name          Type                                    Nulls

ca_row_id            char(15)                                yes
cust_numb            integer                                 no
subr_numb            char(12)                                no
blpd_code            char(2)                                 no
subr_type            char(1)                                 yes
neo_subr_type        char(1)                                 yes
telp_type            char(1)                                 no
titl                 nvarchar(40,10)                         yes
frst_name            nvarchar(80,40)                         no
last_name            nvarchar(80,40)                         yes
subr_stts            char(1)                                 no
from_optr_code       char(3)                                 yes
mnpi_stts            char(1)                                 yes
mnpo_stts            char(1)                                 yes
frst_pkgp_code       char(8)                                 yes
pkgp_strt_date       date                                    yes
new_ppty_code        char(3)                                 yes
dpst_txtp_code       char(4)                                 yes
conx_txtp_code       char(4)                                 yes
telp_allc_type       char(1)                                 yes
hrdw_numb            char(20)                                yes
card_numb            char(19)                                no
sms_lang             char(1)                                 yes
ivr_lang             char(1)                                 yes
ussd_lang            char(1)                                 yes
cc_lang              char(2)                                 yes
deal_numb            char(8)                                 no
slmn_code            char(5)                                 yes
agmt_chck_flag       char(1)                                 yes
agmt_chck_date       date                                    yes
agmt_chck_by         char(12)                                yes
frst_swon_dttm       datetime year to second                 yes
brth_swon_dttm       datetime year to second                 yes
swon_dttm            datetime year to second                 no
swon_by              char(12)                                no
swon_resn_code       char(8)                                 no
swon_area_code       char(1)                                 yes
swof_dttm            datetime year to second                 yes
swof_by              char(12)                                yes
swof_resn_code       char(8)                                 yes
rcnx_resn_code       char(8)                                 yes
leas_code            char(5)                                 yes
leas_strt_date       date                                    yes
leas_expr_date       date                                    yes
leas_acex_date       date                                    yes
pswd_flag            char(1)                                 yes
pswd                 char(4)                                 yes
memb_id              char(20)                                yes
remk                 varchar(255)                            yes
id_type              char(2)                                 yes
id_numb              char(20)                                yes
gndr                 char(1)                                 yes
occp_code            char(3)                                 yes
marl_stts            char(1)                                 yes
salr_levl            char(1)                                 yes
educ_levl            smallint                                yes
emal_addr            varchar(40,0)                           yes
vrfy_emal_dttm       datetime year to second                 yes
wait_emal_addr       varchar(40,0)                           yes
wait_emal_dttm       datetime year to second                 yes
rcnx_paid_flag       char(1)                                 yes
stop_bill_flag       char(1)                                 yes
rfnd_dttm            datetime year to second                 yes
old_cust_numb        integer                                 yes
old_subr_numb        char(12)                                yes
new_cust_numb        integer                                 yes
new_subr_numb        char(12)                                yes
go_inter_flag        char(1)                                 no
chwr_cont_aou        char(1)                                 no
last_chng_dttm       datetime year to second                 no
last_chng_by         char(12)                                no



Index_name         Owner    Type/Clstr Access_Method      Columns

optim_csn_subr02_+ informix dupls/No   btree              telp_type

neo_csp_subr02_00+ informix unique/No  btree              cust_numb
                                                          subr_numb

neo_csf_subr02_00+ informix dupls/No   btree              cust_numb

neo_csf_subr02_00+ informix dupls/No   btree              frst_pkgp_code

neo_csf_subr02_00+ informix dupls/No   btree              swon_resn_code

neo_csf_subr02_00+ informix dupls/No   btree              swof_resn_code

neo_csf_subr02_00+ informix dupls/No   btree              rcnx_resn_code

neo_csf_subr02_00+ informix dupls/No   btree              leas_code

neo_csf_subr02_00+ informix dupls/No   btree              dpst_txtp_code

neo_csf_subr02_00+ informix dupls/No   btree              conx_txtp_code

neo_csf_subr02_00+ informix dupls/No   btree              hrdw_numb

neo_csf_subr02_00+ informix dupls/No   btree              slmn_code

neo_csf_subr02_00+ informix dupls/No   btree              blpd_code
                                                          subr_numb

neo_csf_subr02_00+ informix dupls/No   btree              card_numb

neo_csf_subr02_00+ informix dupls/No   btree              deal_numb

neo_csf_subr02_00+ informix dupls/No   btree              subr_numb

neo_csf_subr02_00+ informix dupls/No   btree              neo_subr_type

neo_csn_subr02_00+ informix dupls/No   btree              swon_dttm

neo_csn_subr02_00+ informix dupls/No   btree              ca_row_id

neo_csn_subr02_00+ informix dupls/No   btree              agmt_chck_date









Column name          Type                                    Nulls

ca_row_id            char(15)                                yes
cust_numb            integer                                 no
blpd_code            char(2)                                 no
comp_code            char(2)                                 no
mgrt_from_cust       integer                                 yes
cust_stts            char(1)                                 no
titl                 varchar(40,10)                          yes
frst_name            nvarchar(80,40)                         no
last_name            nvarchar(80,40)                         yes
id_type              char(2)                                 no
id_numb              char(20)                                no
old_id_numb          char(20)                                yes
gndr                 char(1)                                 yes
date_of_brth         date                                    yes
occp_code            char(3)                                 yes
ntnt_code            char(4)                                 yes
marl_stts            char(1)                                 yes
chld                 smallint                                yes
faml_memb            smallint                                yes
salr_levl            char(1)                                 yes
incl_dict            char(1)                                 yes
good_stts            char(1)                                 yes
lang                 char(1)                                 no
grup_head            char(1)                                 yes
leaf_cust            char(1)                                 yes
pret_cust_numb       integer                                 yes
grup_code            integer                                 no
grup_levl            smallint                                no
empl_flag            char(1)                                 yes
rprt_levl_flag       char(1)                                 no
pmnt_levl_flag       char(1)                                 no
grup_subr_indc       char(1)                                 no
home_telp_numb       char(21)                                yes
home_fax_numb        char(21)                                yes
offc_telp_numb       char(21)                                yes
offc_fax_numb        char(21)                                yes
crcr_telp_numb       varchar(21,0)                           yes
py_telp_numb         char(21)                                yes
emal_addr            varchar(40,0)                           yes
fcbk_id              varchar(40,0)                           yes
fcbk_lgin            varchar(40,0)                           yes
docm_addr_type       char(2)                                 no
acct_type            char(2)                                 yes
acct_sub_type        char(2)                                 yes
jrst_prsn_flag       char(1)                                 yes
brnc_code            varchar(150,0)                          yes
crtd_dttm            datetime year to second                 no
crtd_by              char(12)                                no
last_chng_dttm       datetime year to second                 no
last_chng_by         char(12)                                no



Index_name         Owner    Type/Clstr Access_Method      Columns

csn_cust_07_n      informix dupls/No   btree              id_numb
                                                          cust_stts

csp_cust_01        informix unique/No  btree              cust_numb

csf_cust_04_n      informix dupls/No   btree              occp_code

csf_cust_15_n      informix dupls/No   btree              pret_cust_numb

csn_cust_06_n      informix dupls/No   btree              cust_stts

csn_cust_16_n      informix dupls/No   btree              grup_code

csn_cust_18_n      informix dupls/No   btree              ca_row_id

tsf_cust_02_n      informix dupls/No   btree              blpd_code

tsf_cust_03_n      informix dupls/No   btree              ntnt_code

csf_cust_05_n      informix dupls/No   btree              id_type

csn_cust_13_n      informix dupls/No   btree              frst_name
                                                          blpd_code

csn_cust_14_n      informix dupls/No   btree              last_name

csn_cust_17_n      informix dupls/No   btree              id_numb



ca_row_id       CA_NOT_FOUND
cust_numb       341001902
blpd_code       01
comp_code       10
mgrt_from_cust
cust_stts       C
titl            นาง
frst_name       dtac
last_name       trinet
id_type         15
id_numb         DUMMY
old_id_numb
gndr            M
date_of_brth    1977-01-01
occp_code       505
ntnt_code       66
marl_stts
chld
faml_memb
salr_levl
incl_dict
good_stts       N
lang            T
grup_head       0
leaf_cust       1
pret_cust_numb
grup_code       0
grup_levl       0
empl_flag       N
rprt_levl_flag  N
pmnt_levl_flag  N
grup_subr_indc  0
home_telp_numb  XXX
home_fax_numb   XXX
offc_telp_numb  XXX
offc_fax_numb   XXX
crcr_telp_numb
py_telp_numb
emal_addr       XXX@XXX.com
fcbk_id
fcbk_lgin
docm_addr_type  02
acct_type       -1
acct_sub_type   -1
jrst_prsn_flag
brnc_code
crtd_dttm       1991-10-01 00:00:00
crtd_by         JUPOPER
last_chng_dttm  2018-09-12 17:09:11
last_chng_by    BATCHSB




----------------------- oltp@test_ne51 --------- Press CTRL-W for Help --------

Column name          Type                                    Nulls

seqn                 serial                                  no
cust_numb            integer                                 no
subr_numb            char(12)                                no
pkgp_code            char(8)                                 yes
pack_code            char(8)                                 no
pack_type            char(2)                                 yes
artm_code            smallint                                yes
disc_code            smallint                                yes
pack_strt_dttm       datetime year to second                 no
pack_end_dttm        datetime year to second                 yes
prov_strt_dttm       datetime year to second                 yes
init_end_dttm        datetime year to second                 yes
swof_dttm            datetime year to second                 yes
expr_flag            char(1)                                 no
crtd_dttm            datetime year to second                 no
crtd_by              char(12)                                no
last_chng_dttm       datetime year to second                 no
last_chng_by         char(12)                                no



Index_name         Owner    Type/Clstr Access_Method      Columns

csn_spkd_0001n     informix dupls/No   btree              cust_numb
                                                          subr_numb

csn_spkd_0002n     informix dupls/No   btree              pack_code

csn_spkd_0004n     informix dupls/No   btree              pack_end_dttm
                                                          expr_flag

csp_spkd_0002n     informix unique/No  btree              seqn

csn_spkd_0003n     informix unique/No  btree              cust_numb
                                                          subr_numb
                                                          pack_code
                                                          pack_strt_dttm
                                                          crtd_dttm




seqn            125510010
cust_numb       514996370
subr_numb       66814919760
pkgp_code       12011751
pack_code       12011751
pack_type       10
artm_code       519
disc_code
pack_strt_dttm  2004-08-18 15:19:10
pack_end_dttm   2017-04-20 15:11:03
prov_strt_dttm
init_end_dttm
swof_dttm
expr_flag       1
crtd_dttm       2004-08-18 15:19:10
crtd_by         02223B02
last_chng_dttm  2017-04-20 15:11:03
last_chng_by    DRBTMGRT

seqn            125510011
cust_numb       514996370
subr_numb       66814919760
pkgp_code       12011751
pack_code       11000101
pack_type       10
artm_code       1
disc_code
pack_strt_dttm  2007-08-19 00:00:00
pack_end_dttm   2004-10-18 23:59:59
prov_strt_dttm
init_end_dttm
swof_dttm
expr_flag       1
crtd_dttm       2004-08-18 15:19:10
crtd_by         02223B02
last_chng_dttm  2004-10-19 15:56:56
last_chng_by    SIRIPORE

Column name          Type                                    Nulls

cust_numb            integer                                 no
subr_numb            char(12)                                no
blpd_code            char(2)                                 no
telp_type            char(1)                                 no
swon_dttm            datetime year to second                 no
swof_dttm            datetime year to second                 yes



Index_name         Owner    Type/Clstr Access_Method      Columns

csn_cust_srch_00   informix dupls/No   btree              swon_dttm

csp_cust_srch_00   informix unique/No  btree              cust_numb
                                                          subr_numb

optim_cust_srch_00 informix dupls/No   btree              telp_type



cust_numb  500000600
subr_numb  66816586623
blpd_code  03
telp_type  T
swon_dttm  1998-09-10 17:51:48
swof_dttm  1999-08-10 19:23:25

cust_numb  500000720
subr_numb  66816586955
blpd_code  03
telp_type  T
swon_dttm  1998-09-10 20:13:37
swof_dttm  2001-08-26 16:11:41




Chailuck Kirathisutisathorn

Order Management & EPC unit
CRM & Business solution
IT & Security Division
Technology Group 

M: +66 8 7506 6964
True Corporation Public Company Limited (Head Office)
18 True Tower, RatchadaphisekRoad, HuaiKhwang, 
Bangkok, Thailand 10310
 

