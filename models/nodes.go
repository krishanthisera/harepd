package models

type Slave struct {
	Pid                int64  `gorm:"column:pid;"`
	Status             string `gorm:"column:status;"`
	ReceiveStartLsn    string `gorm:"column:receive_start_lsn;"`
	ReceiveStartTli    int64  `gorm:"column:receive_start_tli;"`
	ReceivedLsn        string `gorm:"column:received_lsn;"`
	ReceivedTli        int64  `gorm:"column:received_tli;"`
	LastMsgSendTime    string `gorm:"column:last_msg_send_time;"`
	LastMsgReceiptRime string `gorm:"column:last_msg_receipt_time;"`
	LatestEndLsn       string `gorm:"column:latest_end_lsn;"`
	LatestEndTime      string `gorm:"column:latest_end_time;"`
	SlotName           string `gorm:"column:slot_name;"`
	SenderHost         string `gorm:"column:sender_host;"`
	SenderPort         string `gorm:"column:sender_port;"`
	Conninfo           string `gorm:"column:conninfo;"`
}

type Master struct {
	Pid             int64  `gorm:"column:pid;"`
	Usesysid        int64  `gorm:"column:usesysid;"`
	Usename         string `gorm:"column:usename;"`
	ApplicationName string `gorm:"column:application_name;"`
	ClientAddr      string `gorm:"column:client_addr;"`
	ClientHostname  string `gorm:"column:client_hostname;"`
	ClientPort      int32  `gorm:"column:client_port;"`
	BackendStart    string `gorm:"column:backend_start;"`
	BackendXmin     string `gorm:"column:backend_xmin;"`
	State           string `gorm:"column:state;"`
	SentLsn         string `gorm:"column:sent_lsn;"`
	WriteLsn        string `gorm:"column:write_lsn;"`
	FlushLsn        string `gorm:"column:flush_lsn;"`
	ReplayLsn       string `gorm:"column:replay_lsn;"`
	WriteLag        string `gorm:"column:write_lag;"`
	FlushLag        string `gorm:"column:flush_lag;"`
	ReplayLag       string `gorm:"column:replay_lag;"`
	SyncPriority    int64  `gorm:"column:sync_priority;"`
	SyncState       string `gorm:"column:sync_state;"`
	ReplyTime       string `gorm:"column:reply_time;"`
}

//type Slave struct {
// 	Pid                int64  `json:"pid"`
// 	Status             string `json:"status"`
// 	ReceiveStartLsn    string `json:"receive_start_lsn"`
// 	ReceiveStartTli    int64  `json:"receive_start_tli"`
// 	ReceivedLsn        string `json:"received_lsn"`
// 	ReceivedTli        int64  `json:"received_tli"`
// 	LastMsgSendTime    string `json:"last_msg_send_time"`
// 	LastMsgReceiptRime string `json:"last_msg_receipt_time"`
// 	LatestEndLsn       string `json:"latest_end_lsn"`
// 	LatestEndTime      string `json:"latest_end_time"`
// 	SlotName           string `json:"slot_name"`
// 	SenderHost         string `json:"sender_host"`
// 	SenderPort         string `json:"sender_port"`
// 	Conninfo           string `json:"conninfo"`
// }

// type Master struct {
// 	Pid 			int64	`json:"pid"`
// 	Usesysid		int64	`json:"usesysid"`
// 	Usename			string	`json:"usename"`
// 	ApplicationName	string	`json:"application_name"`
// 	ClientAddr		string	`json:"client_addr"`
// 	ClientHostname	string	`json:"client_hostname"`
// 	ClientPort		int32	`json:"client_port"`
// 	BackendStart	string	`json:"backend_start"`
// 	BackendXmin		string	`json:"backend_xmin"`
// 	State			string	`json:"state"`
// 	SentLsn			string	`json:"sent_lsn"`
// 	WriteLsn		string	`json:"write_lsn"`
// 	FlushLsn		string	`json:"flush_lsn"`
// 	ReplayLsn		string	`json:"replay_lsn"`
// 	WriteLag		string	`json:"write_lag"`
// 	FlushLag		string	`json:"flush_lag"`
// 	ReplayLag		string	`json:"replay_lag"`
// 	SyncPriority	int64	`json:"sync_priority"`
// 	SyncState		string	`json:"sync_state"`
// 	ReplyTime		string	`json:"reply_time"`

// }
