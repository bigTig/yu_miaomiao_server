package utils

import "time"

func SetCreatedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func SetUpdatedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func setDeletedTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
