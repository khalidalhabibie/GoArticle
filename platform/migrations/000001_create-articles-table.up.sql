BEGIN;


CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

SET TIMEZONE="Asia/Jakarta";

-- Create articles table
CREATE TABLE articles (
    -- id UUID DEFAULT uuid_generate_v4 () PRIMARY KEY,
    id SERIAL PRIMARY KEY,
    author VARCHAR (255) NOT NULL,
    title VARCHAR (255) NOT NULL,
    body  TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL
);

-- Add indexes
CREATE INDEX search_articles ON articles(title, body);
CREATE INDEX search_articles_and_author ON articles(title, body, author);


INSERT INTO "articles" ("author", "title", "body", "created_at") VALUES ('khalid','Bos Baru Chelsea Akan Hadir di Laga Kontra Wolves Usai Menangi Tawaran Rp 71,5 T','Pemilik baru Chelsea, Todd Boehly, diperkirakan akan hadir di tribun penonton untuk menyaksikan klubnya berlaga melawan Wolves di Liga Inggris pada Sabtu (7/5). Konsorsium yang ia pimpin baru saja memenangi tawaran untuk mengakuisisi klub dengan nilai 4,15 miliar pounds atau sekitar Rp 71,5 triliun. \n Proses administrasi dari kesepakatan tersebut telah dilimpahkan ke Pemerintah Inggris dan Premier League. Jika prosesnya berjalan dengan baik, Boehly akan resmi menjadi pemilik baru Chelsea untuk menggantikan Roman Abramovich.Belum genap 24 jam usai kesepakatan fantastis tersebut, Daily Mail melaporkan bahwa Boehly kabarnya akan menghadiri pertandingan Chelsea vs Wolves di Stamford Bridge dalam lanjutan Liga Inggris nanti malam. Boleh jadi, ia akan mendapat sambutan dari fan The Blues.Todd Boehly cukup serius untuk mengakuisisi Chelsea. Pada April lalu, ia juga kedapatan menonton pertandingan Chelsea sekaligus mengunjungi kompleks latihan The Blues. \n Konsorsium Todd Boehly telah ditunjuk sebagai penawar pilihan untuk mengambil alih Chelsea oleh Raine Group, bank Amerika yang mengawasi penjualan. Dia telah bermitra dengan sesama pemilik tim baseball Dodgers, Mark Walter; miliarder Swiss, Hansjorg Wyss; serta perusahaan investasi Clearlake Capital.','2022-05-07 14:31:10.311081');


COMMIT