package util

import (
	"GOKIT_v001/internal/log"
	"database/sql"
	"encoding/json"
	"time"
)

// CREDIT: https://medium.com/aubergine-solutions/how-i-handled-null-possible-values-from-database-rows-in-golang-521fb0ee267

// NullInt64 is an alias for sql.NullInt64 data type
type NullInt64 struct {
	sql.NullInt64
}

// MarshalJSON for NullInt64
func (ni NullInt64) MarshalJSON() ([]byte, error) {
	if !ni.Valid {
		return []byte(`""`), nil
	}
	return json.Marshal(ni.Int64)
}

// NullBool is an alias for sql.NullBool data type
type NullBool struct {
	sql.NullBool
}

// MarshalJSON for NullBool
func (nb NullBool) MarshalJSON() ([]byte, error) {
	if !nb.Valid {
		return []byte(`""`), nil
	}
	return json.Marshal(nb.Bool)
}

// NullFloat64 is an alias for sql.NullFloat64 data type
type NullFloat64 struct {
	sql.NullFloat64
}

// MarshalJSON for NullFloat64
func (nf NullFloat64) MarshalJSON() ([]byte, error) {
	if !nf.Valid {
		return []byte(`""`), nil
	}
	return json.Marshal(nf.Float64)
}

// NullString is an alias for sql.NullString data type
type NullString struct {
	sql.NullString
}

// MarshalJSON for NullString
func (ns NullString) MarshalJSON() ([]byte, error) {
	if !ns.Valid {
		return []byte(`""`), nil
	}
	log.AppTraceLog.Debug(log.AppTraceLogInfo("MARSHALL STRING JSON DATA:"+ns.String, "", "", "", "", ""))
	return json.Marshal(ns.String)
}

// NullTime is an alias for mysql.NullTime data type
type NullTime struct {
	sql.NullTime
}

// MarshalJSON for NullTime
func (nt NullTime) MarshalJSON() ([]byte, error) {
	if !nt.Valid {
		return []byte(`""`), nil
	}
	log.AppTraceLog.Debug(log.AppTraceLogInfo("MARSHALL JSON DATA:"+nt.Time.String(), "", "", "", "", ""))
	return json.Marshal(nt.Time)
}

func (nt *NullTime) UmarshalJSON(data []byte) error {
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		log.AppTraceLog.Debug(log.AppTraceLogInfo("UNMARSHALL JSON DATA:"+nt.Time.String(), "", "", "", "", ""))
		nt.Valid = true
		nt.Time = *x
	} else {
		log.AppTraceLog.Debug(log.AppTraceLogInfo("UNMARSHALL JSON DATA:", "", "", "", "", ""))
		nt.Valid = false
	}
	return nil
}

type OMTime struct {
	time.Time
}

func (t OMTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Time)
}

func (t *OMTime) UnmarshalJSON(b []byte) (err error) {
	date, err := time.Parse(`"2006-01-02T15:04:05"`, string(b))
	if err != nil {
		return err
	}
	t.Time = date
	return
}
