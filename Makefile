migrate_up:
	dbmate -u "postgres://avnadmin:AVNS_Lh14wZ-I2aGddb_u3JM@pg-198915a6-chaitubhojane-46b6.k.aivencloud.com:19568/defaultdb?sslmode=require" up

migrate_down:
	dbmate -u "postgres://avnadmin:AVNS_Lh14wZ-I2aGddb_u3JM@pg-198915a6-chaitubhojane-46b6.k.aivencloud.com:19568/defaultdb?sslmode=require" down