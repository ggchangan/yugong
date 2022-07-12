package model

import (
	"database/sql"
	"time"
)

type ReportMessage struct {
	ObjectMeta `json:"metadata:omitempty"`

	ReportID         sql.NullInt64  `gorm:"column:report_id" json:"-"`
	RunMode          string         `gorm:"column:run_mode" json:"-"`
	Status           string         `gorm:"column:status" json:"status"`
	StartTime        time.Time      `gorm:"column:start_time" json:"start_time"`
	EndTime          *time.Time     `gorm:"column:end_time" json:"end_time"`
	TaskID           sql.NullString `gorm:"column:task_id" json:"-"`
	ReportFileOption sql.NullString `gorm:"column:report_file_option" json:"-"`
	FilePathBase     string         `gorm:"column:file_path_base" json:"-"`
	LocalStoragePath sql.NullString `gorm:"column:local_storage_path" json:"-"`
	RowCount         int64          `gorm:"column:row_count" json:"row_count"`
	FileSize         int64          `gorm:"column:file_size" json:"file_size"`
	Filters          sql.NullString `gorm:"column:filters" json:"filters"`
	ColumnHeader     string         `gorm:"column:column_header" json:"column_header"`
	DateRange        sql.NullString `gorm:"column:date_range" json:"date_range"`
	IsDeleted        bool           `gorm:"column:is_deleted" json:"-"`
	RetryCount       int            `gorm:"column:retry_count" json:"-"`
	CreateUser       string         `gorm:"column:create_user" json:"create_user"`
	NetworkID        int64          `gorm:"column:network_id" json:"network_id"`
	LatencyMinute    int64          `gorm:"column:latency_minute" json:"latency_minute"`
	ErrorDetail      sql.NullString `gorm:"column:error_detail" json:"error_detail"`
	ErrorSummary     sql.NullString `gorm:"column:error_summary" json:"error_summary"`
	Ip               sql.NullString `gorm:"column:ip" json:"ip"`
	NetworkFunction  string         `gorm:"column:network_function" json:"network_function"`
}

// TableName maps to mysql table name.
func (r *ReportMessage) TableName() string {
	return "report_message"
}
