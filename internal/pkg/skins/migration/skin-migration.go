package migration

import "github.com/poportss/jackportcs/internal/migrations"

func Versions() []migrations.Versions {
	return []migrations.Versions{&v1{}, &v2{}, &v3{}, &v4{}, &v5{}}
}
