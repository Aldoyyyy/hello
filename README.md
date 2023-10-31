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