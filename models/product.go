package models

type Product struct {
	id          int64  `gorm:"primaryKey" json:"id"`
	namaProduct string `gorm:"type:varchar(300)" json:"nama_product"`
	deskripsi   string `gorm:"type:text" json:"deskripsi"`
}
