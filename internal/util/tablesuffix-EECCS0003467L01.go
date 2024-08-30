package util

import (
	"GOKIT_v001/internal/log"
	"errors"
	"strconv"
	"strings"
)

const (
	TAB_ERROR          = "--NO TABLE--"
	TAB_NAME_SPKD      = "CS_SPKD"
	TAB_NAME_SUBR      = "CS_SUBR"
	TAB_NAME_CUST_SRCH = "CS_CUST_SRCH"
	TAB_NAME_ADD_SUBR  = "CS_ADD_SUBR"

	PROD_PCN = "PCN"
	PROD_IEC = "IEC"

	INDC_PROD_2SUBR = "INDC_PROD_2SUBR"
	INDC_PROD_2CUST = "INDC_PROD_2CUSt"
	INDC_2CUST      = "INDC_2CUSt"
)

var tab_mod = map[string]int{
	TAB_NAME_SPKD:      10,
	TAB_NAME_SUBR:      10,
	TAB_NAME_CUST_SRCH: 20,
	TAB_NAME_ADD_SUBR:  10,
}

var tab_indc = map[string]string{
	TAB_NAME_SPKD:      INDC_PROD_2SUBR,
	TAB_NAME_SUBR:      INDC_PROD_2SUBR,
	TAB_NAME_CUST_SRCH: INDC_2CUST,
	TAB_NAME_ADD_SUBR:  INDC_PROD_2SUBR,
}

func getSuffixText(indc_numb string, tabl_name string, digit int, lt log.LogTracing) (string, error) {

	i_subr_numb, err := strconv.Atoi(indc_numb)

	if err != nil {
		log.AppTraceLog.Error(log.GenErrLog("Subscriber is not integer - SUBR:"+indc_numb, lt, log.E100010, nil))
		return "", errors.New("cannot convert subscriber number to string")
	}

	suff := strconv.Itoa(i_subr_numb % tab_mod[tabl_name])

	if len(suff) < digit {
		gap := digit - len(suff)

		zeroText := strings.Repeat("0", gap)
		suff = zeroText + suff
	}
	return suff, nil
}

func GetTableSuffix(cust_numb string, subr_numb string, tabl_name string, prod_type string, sub_type string, lt log.LogTracing) (string, string, error) {
	isNotEmpty := IsNotEmptyString(tabl_name, prod_type)

	log.AppTraceLog.Debug(log.AppTraceLogInfo("Begin GetTable Suffix:"+tabl_name, "", "", subr_numb, "", ""))

	if !isNotEmpty {
		err := errors.New("string is empty")

		log.AppTraceLog.Error(log.GenErrLog("Blank is not allow: table_name, prod_type", lt, log.E100009, nil))

		return TAB_ERROR, TAB_ERROR, err
	}
	ret_tab_name := tabl_name

	log.AppTraceLog.Debug(log.AppTraceLogInfo("Mode:"+tab_indc[tabl_name], "", "", "", "", ""))
	if tab_indc[tabl_name] == INDC_PROD_2SUBR {

		suffix, err := getSuffixText(subr_numb, tabl_name, 2, lt)
		if err != nil {
			log.AppTraceLog.Error(log.GenErrLog("Cannot find suffix of table: "+tabl_name+"SUBR:"+subr_numb, lt, log.E100009, nil))

			return TAB_ERROR, TAB_ERROR, nil
		}
		ret_tab_name = ret_tab_name + "_" + prod_type + "_" + suffix
		return ret_tab_name, suffix, nil
	}

	if tab_indc[tabl_name] == INDC_PROD_2CUST {
		suffix, err := getSuffixText(cust_numb, tabl_name, 2, lt)
		if err != nil {
			log.AppTraceLog.Error(log.GenErrLog("Cannot find suffix of table: "+tabl_name+"SUBR:"+subr_numb, lt, log.E100009, nil))

			return TAB_ERROR, TAB_ERROR, err
		}
		ret_tab_name = ret_tab_name + "_" + prod_type + "_" + suffix
		return ret_tab_name, suffix, nil
	}

	if tab_indc[tabl_name] == INDC_2CUST {
		suffix, err := getSuffixText(cust_numb, tabl_name, 2, lt)
		if err != nil {
			log.AppTraceLog.Error(log.GenErrLog("Cannot find suffix of table: "+tabl_name+"SUBR:"+subr_numb, lt, log.E100009, nil))

			return TAB_ERROR, TAB_ERROR, err
		}
		ret_tab_name = ret_tab_name + "_" + suffix
		return ret_tab_name, suffix, nil
	}

	return TAB_ERROR, TAB_ERROR, errors.New("unexpected error")
}
