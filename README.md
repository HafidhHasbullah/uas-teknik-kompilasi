# Simulasi Compiler: Konstruksi Perulangan `while`

Repositori ini berisi laporan dan implementasi Tugas Ujian Akhir Semester untuk mata kuliah Compiler. Proyek ini menyimulasikan tahapan kompilasi pada konstruksi sintaksis perulangan `while`.

## Informasi Mahasiswa

| Deskripsi | Keterangan |
| :--- | :--- |
| **Nama Lengkap** | Hafidh Hasbullah |
| **Nomor Induk Mahasiswa (NIM)** | 231011401315 |
| **Kelas** | 06TPLE006 |

---

## 1. Pilihan Konstruksi
Konstruksi sintaksis yang dipilih untuk disimulasikan pada tugas ini adalah **Perulangan (*Looping*) menggunakan `while`**.

## 2. Pola Tata Bahasa (*Grammar* / BNF)
Aturan sintaksis untuk konstruksi perulangan `while` ini didefinisikan menggunakan pendekatan *Backus-Naur Form* (BNF) sederhana berikut:

```text
<while_stmt> ::= "while" "(" <condition> ")" "{" <statements> "}"
<condition>  ::= <identifier> <operator> <value>
<statements> ::= <identifier> "=" <expression>
<expression> ::= <identifier> | <identifier> <operator> <value>
<identifier> ::= [A-Za-z_][A-Za-z0-9_]*
<value>      ::= [0-9]+
<operator>   ::= "<" | ">" | "==" | "+" | "-" | "*" | "/"

```

## 3. Penjelasan Implementasi Tahapan Kompilasi

Program ini menerima input *source code* berupa *string* (contoh: `while ( x > 0 ) { x = x - 1 }`) dan memprosesnya secara berurutan melalui 4 tahapan utama:

### A. Analisis Leksikal (*Lexical Analysis*)

Pada tahap ini, program memecah *string* input mentah menjadi kumpulan token yang bermakna.

* **Metode:** Menggunakan *Regular Expression* (Regex) untuk mencocokkan dan mengenali pola karakter.
* **Proses:** Program mengkategorikan potongan kode menjadi token seperti *keyword* (`while`), tanda kurung (`(`, `)`), kurung kurawal (`{`, `}`), operator (`>`, `-`), *identifier* (variabel seperti `x`), dan angka (`0`, `1`). Spasi kosong (*whitespace*) akan diabaikan, sedangkan karakter yang tidak lazim akan memicu *error*.

### B. Analisis Sintaksis (*Syntax Analysis*)

Tahap ini memvalidasi urutan token untuk memastikan susunannya sesuai dengan tata bahasa (BNF) yang ditetapkan.

* **Metode:** Mengonstruksi *Abstract Syntax Tree* (AST) sederhana berbasis struktur data *Map* atau *Dictionary*.
* **Proses:** Program memastikan urutan token diawali dengan `while`, diikuti oleh kurung buka, kondisi, kurung tutup, dan isi (*body*) di dalam kurung kurawal. Jika urutan salah atau tidak lengkap, program akan menggagalkan kompilasi dan menampilkan `SyntaxError`.

### C. Analisis Semantik (*Semantic Analysis*)

Tahap ini melakukan pengecekan makna dan validitas variabel yang digunakan dalam *source code*.

* **Metode:** Pengecekan silang terhadap *Symbol Table* (Tabel Simbol) tiruan.
* **Proses:** Program memeriksa setiap *identifier* (variabel) di dalam kondisi maupun *body* perulangan. Jika ada variabel yang digunakan tetapi belum dideklarasikan di dalam *Symbol Table*, program mendeteksinya sebagai kesalahan semantik dan menampilkan `NameError`.

### D. Generasi Kode Antara (*Three-Address Code* / TAC)

Tahap terakhir ini bertugas menerjemahkan AST yang sudah tervalidasi menjadi instruksi tingkat menengah (*Three-Address Code*).

* **Proses:**
1. Men-*generate* label unik untuk awal perulangan (contoh: `L1`) dan akhir perulangan (`L2`).
2. Mengevaluasi kondisi menggunakan lompatan bersyarat: `ifFalse [kondisi] goto L2` (jika kondisi salah, lompat ke akhir).
3. Mengurai *body* perulangan. Jika terdapat operasi aritmatika (misalnya `x - 1`), program membuat variabel *temporary* sementara (misal `t1`) untuk menampung hasil operasi sebelum disimpan ke variabel target.
4. Menambahkan instruksi `goto L1` di akhir *body* untuk mengarahkan eksekusi kembali ke awal *loop*.



```

```
