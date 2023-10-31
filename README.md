# hello
Finding Reference:
1. Development method: SCRUM.
```
metodologi yang termasuk dalam agile software development. Scrum dinilai dapat menghasilkan kualitas perangkat lunak yang baik sesuai dengan keinginan pengguna, dapat digunakan dalam proyek besar maupun kecil, dan mudah untuk mengadopsi perubahan.
```
2. Monolithic VS Microservice.
```
Monolitik adalah sebuah arsitektur aplikasi secara kesatuan atau tunggal, maka microservices adalah sebaliknya. Microservices terbagi menjadi unit pecahan yang lebih kecil dan spesifik.
```
3. Array vs Map
```
Array Adalah : Pengertian, Kegunaan, dan Jenisnya – Apa itu array? Pastinya kata Array tidak asing didengar pada bidang pemprograman. Dengan array akan memudahkan dalam penyimpanan data banyak yang sudah dikelompokkan. Array terdapat dalam Phyton, C++, PHP, Java dan bahasa pemrograman lainnya.

Map adalah tipe data asosiatif yang ada di Go, berbentuk key-value pair. Untuk setiap data (atau value) yang disimpan, disiapkan juga key-nya. Key harus unik, karena digunakan sebagai penanda (atau identifier) untuk pengaksesan value yang bersangkutan.
```
4. Json
```
Json adalah sebuah format untuk berbagi data. Seperti dapat kita lihat dari namanya, JSON diturunkan dari bahasa pemrograman javaScript, akan tetapi format ini tersedia bagi banyak bahasa lain termasuk Python, Ruby, PHP, dan Java. JSON biasanya dilafalkan seperti nama "Jason.
Contoh :
{
  "first_name" : "Sammy",
  "last_name" : "Shark",
  "location" : "Ocean",
  "online" : true,
  "followers" : 987 
}
```

Code:
1. Programming Language: Go
2. Database: Mysql
3. API: REST
## Tutorial Run Go
1. `go run main.go`

## Tutorial Git Transfer
### First Git Project
1. `git status`
```
On branch master

No commits yet  

Untracked files:
  (use "git add <file>..." to include in what will be committed)
        README.md
        go.mod
        main.go

nothing added to commit but untracked files present (use "git add" to track)
```
2. `git add .`
3. `git status`
```
On branch master

No commits yet

Changes to be committed:
  (use "git rm --cached <file>..." to unstage)
        new file:   README.md
        new file:   go.mod
        new file:   main.go
```
4. `git commit -m "first commit"`
```
Author identity unknown

*** Please tell me who you are.

Run

  git config --global user.email "you@example.com"
  git config --global user.name "Your Name"

to set your account's default identity.
Omit --global to set the identity only in this repository.

fatal: unable to auto-detect email address (got 'User@WINDOWS-RLFGAUE.(none)')
```
5. `git config --global user.email "stevenalvonso@gmail.com"`
6. `git config --global user.name "Aldoyyyy"`
7. `git commit -m "first commit"`
```
[master (root-commit) 98d86a9] first commit
 3 files changed, 19 insertions(+)
 create mode 100644 README.md
 create mode 100644 go.mod
 create mode 100644 main.go
```
8. `git remote add origin https://github.com/Aldoyyyy/hello.git`
9. `git push -u origin master`
```
info: please complete authentication in your browser...
Enumerating objects: 5, done.
Counting objects: 100% (5/5), done.
Delta compression using up to 12 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (5/5), 537 bytes | 537.00 KiB/s, done.
Total 5 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/Aldoyyyy/hello.git
 * [new branch]      master -> master
branch 'master' set up to track 'origin/master'.
```
### Update Git Project
1. `git status`
```
On branch master
Your branch is up to date with 'origin/master'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md
        modified:   main.go

no changes added to commit (use "git add" and/or "git commit -a")
```
2. `git add .`
3. `git status`
```
On branch master
Your branch is up to date with 'origin/master'.

Changes to be committed:
  (use "git restore --staged <file>..." to unstage)
        modified:   README.md
        modified:   main.go
```
4. `git commit -m "add new function IsMoreThanFive"`
```
[master 479f4fb] add new function IsMoreThanFive
 2 files changed, 83 insertions(+), 1 deletion(-)
```
5. `git push -u origin master`
```
Enumerating objects: 7, done.
Counting objects: 100% (7/7), done.
Delta compression using up to 12 threads
Compressing objects: 100% (4/4), done.
Writing objects: 100% (4/4), 1.41 KiB | 1.41 MiB/s, done.
Total 4 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/Aldoyyyy/hello.git
   98d86a9..479f4fb  master -> master
branch 'master' set up to track 'origin/master'.
```

## Tutorial git clone (copy project github)
1. `git clone https://github.com/Aldoyyyy/hello.git`
```
Cloning into 'hello'...
remote: Enumerating objects: 22, done.
remote: Counting objects: 100% (22/22), done.
remote: Compressing objects: 100% (17/17), done.
remote: Total 22 (delta 7), reused 19 (delta 4), pack-reused 0
Receiving objects: 100% (22/22), 4.07 KiB | 4.07 MiB/s, done.
Resolving deltas: 100% (7/7), done.
```