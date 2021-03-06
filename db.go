package ssrpanel

import (
	"time"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	ID      uint
	VmessID string `gorm:"column:v2ray_uuid"`
	Email   string `gorm:"column:email"`
	Port    int
}

func (*UserModel) TableName() string {
	return "user"
}

type UserTrafficLog struct {
	ID       uint `gorm:"primary_key"`
	UserID   uint
	Uplink   uint64 `gorm:"column:u"`
	Downlink uint64 `gorm:"column:d"`
	NodeID   uint
	Rate     float64
	Traffic  string
	LogTime  int64
}

func (l *UserTrafficLog) BeforeCreate(scope *gorm.Scope) error {
	l.LogTime = time.Now().Unix()
	return nil
}

type NodeOnlineLog struct {
	ID         uint `gorm:"primary_key"`
	NodeID     uint
	OnlineUser int
	LogTime    int64
}

func (*NodeOnlineLog) TableName() string {
	return "ss_node_online_log"
}

func (l *NodeOnlineLog) BeforeCreate(scope *gorm.Scope) error {
	l.LogTime = time.Now().Unix()
	return nil
}

type NodeIP struct {
	ID        uint `gorm:"primary_key"`
	NodeID    uint
	UserID    uint
	Port      int
	IPList    string `gorm:"column:ip"`
	CreatedAt int64
}

func (*NodeIP) TableName() string {
	return "ss_node_ip"
}

func (n *NodeIP) BeforeCreate(scope *gorm.Scope) error {
	n.CreatedAt = time.Now().Unix()
	return nil
}

type NodeInfo struct {
	ID      uint `gorm:"primary_key"`
	NodeID  uint
	Uptime  time.Duration
	Load    string
	LogTime int64
}

func (*NodeInfo) TableName() string {
	return "ss_node_info_log"
}

func (l *NodeInfo) BeforeCreate(scope *gorm.Scope) error {
	l.LogTime = time.Now().Unix()
	return nil
}

type Node struct {
	ID          uint `gorm:"primary_key"`
	TrafficRate float64
}

func (*Node) TableName() string {
	return "ss_node"
}

type DB struct {
	DB         *gorm.DB
	RetryTimes int64
}

func (db *DB) GetAllUsers(nodeClass string) ([]UserModel, error) {
	users := make([]UserModel, 0)
	err := db.DB.Select("id, v2ray_uuid, email, port").Where("enable = 1 AND u + d < transfer_enable AND plan >= ?", nodeClass).Find(&users).Error
	return users, err
}

func (db *DB) GetNode(id uint) (*Node, error) {
	node := Node{}
	err := db.DB.First(&node, id).Error
	return &node, err
}
