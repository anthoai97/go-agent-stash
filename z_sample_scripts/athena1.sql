/*  This is an Athena example to review logs data
time
method
status_code
domain
path
host_ip
bytes_out
response_processing_time	
api_key_id
*/

-- 2023-07-18T10:05:01.976Z GET 200 localhost /organization/v0/properties 0:0:0:0:0:0:0:1 4264 348 wUN0NeTMByHVmIOjo79ImwylAjjkRtkr

-- Athena query processing
CREATE EXTERNAL TABLE IF NOT EXISTS custom_logs (
	time string,
	method string,
	status_code int,
	domain string,
	path string,
	host_ip string,
	bytes_out bigint,
	response_processing_time double,
	api_key_id string
	)
	PARTITIONED BY
	(
		day STRING
	)
	ROW FORMAT DELIMITED  
		FIELDS TERMINATED BY ' '
		LINES TERMINATED BY '\n'
	LOCATION '<s3://path>/'
	TBLPROPERTIES
	(
		"projection.enabled" = "true",
		"projection.day.type" = "date",
		"projection.day.range" = "2022/01/01,NOW",
		"projection.day.format" = "yyyy/MM/dd",
		"projection.day.interval" = "1",
		"projection.day.interval.unit" = "DAYS",
		"storage.location.template" = "<s3://path>/${day}"
	)