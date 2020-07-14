package repository

const (
	aeroUserNamespace       = "user"
	aeroRealTimeDataSetName = "real_time_data"
	aeroUserDataBin         = "user_data" // Data Format [min_magic_day, max_magic_day, updated_ts[YYYYDDMM]]    If this data is not there then fetch the data from redShift
	aeroEntityDataBin       = "_id"
)
