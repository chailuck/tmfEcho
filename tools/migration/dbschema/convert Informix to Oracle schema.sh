########### cut here ###############
#!/bin/bash
# filename: conversion_config.bsh
# Configuration for conversion.

# List of tables to convert, can be lots or few.

TABLE_LIST='
A_table_name
B_table_name'

AWK="/usr/bin/gawk"
SCHEMA_FILE="schema.sql"
TABLE_DATA_PATH="table_data" # Directory for table data.
mkdir $TABLE_DATA_PATH
TABLE_CONTROL_FILE_PATH="control_file" # Oracle sqlldr control files.
mkdir $TABLE_CONTROL_FILE_PATH
LOAD_LOG_PATH="load_log" # Table loading logs.
mkdir $LOAD_LOG_PATH

########### cut here ###############

#!/bin/bash
# filename: create_schema.bsh
# Run the Oracle schema script, load tables with Oracles loader
# Note: make sure Oracle environment variables are set.

# Config
. ./conversion_config.bsh
LOADER="sqlldr"
# For Informix to Oracle date formats:
export NLS_LANG='american_america.us7ascii' # Oracle juju for date format.
export NLS_DATE_FORMAT='MM/DD/YYYY' # Set Oracle to USA Informix date format.

# Create Oracle tables
sqlplus << EOF
user_name/password
@${SCHEMA_FILE}

EOF

# Load tables into Oracle
# $TABLE_CONTROL_FILE_PATH/$table.bad is the file with bad records.
for table in $TABLE_LIST
do
   ${LOADER} control=${TABLE_CONTROL_FILE_PATH}/$table.ctl log=${LOAD_LOG_PATH}
done

#################### cut here ######################################

#!/bin/bash
# filename: get_schema.bsh
# Get Informix schema, convert to Oracle schema DDL.
# Then
# unload Informix tables to files named .txt
# Note: make sure Informix environment variables are set.

. ./I-O_config.bsh

# Get Informix schema DDL.
get_schema () {
for table in ${TABLE_LIST}
do
   echo "drop table $table;"
   dbschema -d order_entry -t $table
done
}

# Remove extra text from Informix schema.
remove_dbschema_header () {
grep -v 'DBSCHEMA Schema Utility'|
grep -v 'Copyright (C) Informix Software'|
grep -v 'Software Serial Number'|
grep -v '{ TABLE'|
grep -v ' }'
}

# Convert Informix datatypes to Oracle
convert_datatypes () {
${AWK} '
/ serial/ {gsub(" serial", " number")}
/ money/ {gsub(" money", " number")}
/ integer/ {gsub(" integer", " number")}
/ decimal/ {gsub(" decimal", " number")}
/ smallint/ {gsub(" smallint", " number")}
/ char/ {gsub(" char", " varchar2")}
/ informix/ {gsub("\"informix\".", "")} # Remove user from DDL.
/ revoke all/ {next}  # Skip permission granting.
{print}'
}

get_table_columns () {

dbaccess database_name << EOF | grep -v '^$'
   output to pipe "cat" without headings
   select colname from syscolumns, systables
   where
   systables.tabname = "$table"
   and systables.tabid = syscolumns.tabid;
   -- order by colno; May use if columns are NOT in correct order.

EOF
}

# Informix unload.
unload_tables () {

for table in ${TABLE_LIST}
do
dbaccess database_name << EOF

unload to "${TABLE_DATA_PATH}/${table}/.out"
select * from $table;

EOF
done
}

# Create Oracle control files. 
make_control_file () {
for table in ${TABLE_LIST}
do
cat << EOF > $TABLE_CONTROL_FILE_PATH/$table.ctl

load data
infile '${TABLE_DATA_PATH}/${table}/.out'
into table $table
fields terminated by "|"
EOF

  echo '(' >> $TABLE_CONTROL_FILE_PATH/$table.ctl
  COLUMNS=$(get_table_columns)
  echo $COLUMNS |
  ${AWK} '{ gsub (" ",","); print}' >> $TABLE_CONTROL_FILE_PATH/$table.ctl
  echo ')' >> $TABLE_CONTROL_FILE_PATH/$table.ctl
done
}

#################
# Main

get_table_schema |
remove_dbschema_header |
convert_datatypes > $ORACLE_SCHEMA

unload_tables

make_control_file


#############################################################################