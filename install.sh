 llgo build -target esp32 -v -tags nogc,emb -o t.bin
 python -m esptool --chip esp32 -b 115200 --before default_reset --after hard_reset write_flash --flash_mode dio --flash_size 2MB --flash_freq 40m 0x1000 t.bin
