/*  This is an Athena example to review logs data
time
status_code
domain
path
host_ip
bytes_in
bytes_out
response_processing_time	
api_key_id
*/

-- 2023-7-17T22:23:00.186641Z 200 api.gauvendi.com /inventory/v1/product 127.0.0.12 12300 12356 2000 F61f2DvRvZ1bvPXdAS3PggCpCaWGj1uD

-- Athena query processing
CREATE EXTERNAL TABLE IF NOT EXISTS custom_logs (
	time string,
	status_code int,
	domain string,
	path string,
	host_ip string,
	bytes_in bigint,
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