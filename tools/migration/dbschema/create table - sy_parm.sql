--Oracle

DROP TABLE SY_PARM;
create table SY_PARM (
    card_numb varchar(2) NOT NULL,
    hrdw_stts varchar(2) NOT NULL,
    telp_type varchar(3) NOT NULL,
    tnst_time int NOT NULL,
    slot_telp_type varchar(1) NOT NULL,
    last_etrc_besn DATE NOT NULL,
    init_year TIMESTAMP(0) NOT NULL,
    vers_numb varchar(10),
    rels_date DATE ,
    pswd_expr_perd int ,
    max_logn int 
);