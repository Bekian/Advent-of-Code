package main

import f "core:fmt"
import os "core:os"

main :: proc() {
	buffer: []byte
	handle := os.open("input.txt", 0)
	read = os.read(handle, buffer)
	f.print(read)
}
