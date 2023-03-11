
# Concurrent Programming

Pada session 12 mempelajari tentang sequential program, paralel program, concurrent program, conccurrency allow parallelism, goroutines, gomaxprocs, channel, select, race condition, blocking with waitgroups 


## Resume

 - Sequential program merepresentasikan suatu program yang menjalankan tugas masing-masing sesuai urutan dalam penyelesaiannya.

 - Parallel program merepresentasikan suatu program yang mampu menjalankan semua tugas bersama-sama dengan waktu yang bersamaan.

- Concurrent program merepresentasikan suatu program yang menjalankan tugas mandiri atau independen dengan waktu yang bersamaan. 

- Conccurrency allow parallelism merepresentasikan concurrency goroutine yang mampu mempermudah pembuatan program paralel dengan prosesor yang disediakan.

- Goroutine mempunyai beberapa fitur yang tersedia seperti synchronization and messaging (channels) yang digunakan untuk berkomunikasi satu sama lain dan multi-way concurrent control (select) yang digunakan untuk mengontrol komunikasi data melalui satu atau banyak channel

- Race condition yang merepresentasikan sebagai program yang mengalami dua kondisi mengakses memori atau menulis dengan membaca memori. kondisi balapan terjadi karena akses yang tidak sinkron ke memori bersama.

- Blocking with waitgroups digunakan sebagai penyelesaian data race dengan memblokir akses membaca memori hingga operasi menulis selesai dijalankan



## Documentation

[Documentasi Praktikum](https://github.com/nadiahindrianti/go_nadiah-indrianti/tree/main/12_Concurrent%20Programming/Screenshot)


## Used By

Nadiah Indrianti

