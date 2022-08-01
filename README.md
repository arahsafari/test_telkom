# test_telkom
## Prerequisites

- Golang
- `make` command
- docker-compose
- docker

## Jawaban
#### Soal1
1. Di Indonesia, ada pecahan mata uang rupiah, yaitu :
* 100.000,-
* 50.000,-
* 20.000,-
* 10.000,-
* 5.000,-
* 2.000,-
* 1.000,-
* 500,-
* 200,-
* 100,-
  Buatlah sebuah fungsi untuk menghitung berapa lembar pecahan yang harus dikeluarkan dari input harga (dengan pembulatan ke atas jita punya harga pecahan antara 1 sampai 99)
  Input : 145.000
  Output:
  {
  “Rp. 100.000”:1,
  “Rp. 20.000”:2,
  “Rp. 5.000”:1,
  }
  Input: 2050
  Ouput:
  {
  “Rp. 2.000”:1,
  “Rp. 100”:1,
  }

```shell script
$ make soal1
```

#### Soal2
2. Anda diminta untuk membuat sebuah function dimana function tersebut berfungsi untuk menentukan apakah dari dua data string yang diberikan membutuhkan sekali proses edit atau lebih. Jika lebih dari sekali proses edit berarti function tersebut akan mengembalikan response False, sedangkan jika hanya sekali proses edit maka function tersebut akan mengembalikan response True. Proses edit di sini dapat berarti melakukan insert sebuah character, remove sebuah character, atau replace sebuah character.

```shell script
$ make soal2
```

#### Soal3
3. Apakah ada kesalahan dari script di bawah ini? Jika ada tolong jelaskan dimana letak kesalahannya dan bagaimana anda memperbaikinya. Jika tidak ada, tolong jelaskan untuk apa script di bawah ini.

Jawaban :
- Penjelasan jawaban ada di Dockerfile.fix yah , terima kasih

#### Soal4
Menurut anda apakah tujuan penggunaan microservices?

Jawaban:

- tujuan microservices biasanya agar mempermudah tim dalam mempelajari service baru karena scopenya
yang di perkecil sampai level domain dan dalam pengerjaan complex architecture menjadi lebih mudah di breakdown kepada tim-tim yang lebih kecil , berbeda
dengan monolith ketika system tersebut udah jadi dan ada newcomer maka newcomer tersebut
akan kesulitan untuk membaca system yang sebesar itu.
- microservices juga mempersingkat waktu jiga ada major refactor code , sehingga
membuat microservice bisa dibuat seflexible mungkin , berbeda dengan monolith
yang ketika ingin refactor biasanya memakan waktu yang sangat lama
- microservice juga dapat membuat servicenya tidak dependant / loosely coupled terhadap service lain
dimana service tersebut tidak akan mengalami failure jika ada failure di salah satu
node service tersebut
- database microservice dapat di sesuaikan terhadap data yang akan kita olah terhadap service tersebut



#### Soal5
5. Buatlah Sebuah backend api dengan golang untuk melakukan health check kepada service backend

step 1
```shell script
$ make soal5
```

step 2
```shell script
$ docker ps -a
```
untuk melihat status health dari service tersebut

