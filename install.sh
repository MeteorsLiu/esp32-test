 python /Users/haolan/esp/esp-idf/components/esptool_py/esptool/esptool.py --chip esp32 elf2image --flash_mode dio --flash_freq 40m --flash_size 2MB --min-rev-full 0 --max-rev-full 399 -o t.bin t.elf
 python -m esptool --chip=esp32 --port /dev/cu.usbserial-10 write_flash 0x1000 t.bin -ff 80m -fm dout
