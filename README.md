# belajarbanyak
pertemuan 13 mei 2026 mppl


https://webhook.cool/at/polite-lion-17/FTr4gP9_ueZmb_LW5OkWUfnopMk0FdCK 


https://polite-lion-17.webhook.cool 




UNTUK UPDATE LOCAL FOLDER 'belajarbanyak' (lakukan ini setiap mau push)

```
cd belajarbanyak

git pull
```

JIKA MAU UPDATE MAKA

```
1. copy file ke folder `belajarbanyak`
2. ketik...

git pull

git add .

git commit -m "Pesan disini"

git push
```

selesai... jika ada masalah pada git pull maka sudah ada yg terbaru... harusnya jangan pindahin file baru ke 'banyakbelajar' dulu... tapi bisa dengan..

```
1. Ctrl+C untuk membatalkan file
2. ketik :qa
3. ketik....

git merge main -m "Pesan disini"

// dari situ akan gabung (merge) dari github yang terbaru + file anda yang baru

```


CARA SETUP AWAL GIT

```
git clone https://github.com/adoobi/belajarbanyak

cd belajarbanyak

git config --global user.email "masukkan email anda"

git pull

git push
```




CARA SETUP DARI AWAL

```
git init

git remote add origin https://github.com/adoobi/belajarbanyak

git config --global user.email "masukkan email anda"

// kalau udah ada...

git fetch --all

git pull

git branch -a   // liat all branch

git push -u origin main
```
