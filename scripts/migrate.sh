#!/bin/bash

# Migrasyon komutlarının bulunduğu dizin yolu
MIGRATIONS_DIR="/path/to/migrations"

# Veritabanı bağlantı URL'si
DB_URL="postgresql://user:password@localhost:5432/database_name"

# migrate komutunun bulunduğu dizini doğrula
if ! command -v migrate &> /dev/null
then
    echo "migrate komutu bulunamadı. Lütfen migrate'ı yükleyin ve PATH'e ekleyin."
    exit 1
fi

# Migrasyonları uygula
migrate -database "${DB_URL}" -path "${MIGRATIONS_DIR}" up

# Migrasyon komutunun başarılı olup olmadığını kontrol et
if [ $? -ne 0 ]; then
    echo "Migrasyon başarısız oldu."
    exit 1
else
    echo "Migrasyon başarıyla tamamlandı."
fi
