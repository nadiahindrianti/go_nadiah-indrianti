
# ORM & Code Stucture (MVC)

Pada bab ini mempelajari tentang Database Connection & migration, CRUD Restfull, MVC.

## Resume

ORM dalam ilmu komputer adalah teknik pemrograman untuk mengubah data antara sistem tipe yang tidak kompatibel menggunakan bahasa pemrograman berorientasi objek

Database migration merupakan cara memperbarui versi basis data untuk menyesuaikan perubahan versi aplikasi dengan upgrade ke versi terbaru atau rollback ke versi sebelumnya.

Untuk mengoneksikan code dengan database dengan menggunakan func InitDB serta InitMigration.
## Running

Install

```
  go get -u gorm.io/gorm
  go get -u gorm.io/driver/mysql
```


Connection Database

```
  dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
```


Contoh All Rest API routes

```
  e.GET("/users/:id", GetUserController)
  e.DELETE("/users/:id", DeleteUserController)
  e.PUT("/users/:id", UpdateUserControler)
```
## Documentation

[Documentasi Praktikum](https://github.com/nadiahindrianti/go_nadiah-indrianti/tree/main/20_ORM%20%26%20Code%20Structure%20(MVC)/Screenshot)


## Used By

Nadiah Indrianti

