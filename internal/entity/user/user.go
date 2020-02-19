package user

import "time"

// User object model
type User struct {
	ID       int       `db:"id" json:"user_id"`
	Nama     string    `db:"nama" json:"nama"`
	Umur     int       `db:"umur" json:"umur"`
	Alamat   string    `db:"alamat" json:"alamat"`
	TglLahir time.Time `db:"tanggal_lahir" json:"tgl_lahir"`
}
