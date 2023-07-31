package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Members struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	M_Nama  string             `bson:"m_nama,omitempty" json:"m_nama,omitempty"`
	M_Study string             `bson:"m_study,omitempty" json:"m_study,omitempty"`
}

type Customers struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	C_Nama  string             `bson:"c_nama,omitempty" json:"c_nama,omitempty"`
	C_Study string             `bson:"c_study,omitempty" json:"c_study,omitempty"`
}

type Incomes struct {
	ID      primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Qty     int                `bson:"qty,omitempty" json:"qty,omitempty"`
	Halaman int                `bson:"halaman,omitempty" json:"halaman,omitempty"`
	Uang    int                `bson:"uang,omitempty" json:"uang,omitempty"`
}

type Admin struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	Username string             `bson:"username,omitempty" json:"username,omitempty"`
	Password string             `bson:"password,omitempty" json:"password,omitempty"`
}

type Token struct {
	Token_String string `bson:"tokenstring,omitempty" json:"tokenstring,omitempty"`
}
