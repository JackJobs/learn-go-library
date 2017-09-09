package main

import (
	"encoding/binary"
	"fmt"
	"os"
)

/*
BITMAPCOREHEADER structure:
typedef struct tagBITMAPCOREHEADER {
  DWORD bcSize;
  WORD  bcWidth;
  WORD  bcHeight;
  WORD  bcPlanes;
  WORD  bcBitCount;
} BITMAPCOREHEADER, *PBITMAPCOREHEADER;

BITMAPINFOHEADER structure:
typedef struct tagBITMAPINFOHEADER {
  DWORD biSize;
  LONG  biWidth;
  LONG  biHeight;
  WORD  biPlanes;
  WORD  biBitCount;
  DWORD biCompression;
  DWORD biSizeImage;
  LONG  biXPelsPerMeter;
  LONG  biYPelsPerMeter;
  DWORD biClrUsed;
  DWORD biClrImportant;
} BITMAPINFOHEADER, *PBITMAPINFOHEADER;

*/

type BitmapInfoHeader struct {
	Size           uint32
	Width          int32
	Height         int32
	Places         uint16
	BitCount       uint16
	Compression    uint32
	SizeImage      uint32
	XPerlsPerMeter int32
	YPerlsPerMeter int32
	ClsUsed        uint32
	ClrImportant   uint32
}

//read bmp file header info
func main() {
	file, err := os.Open("io_bmp.bmp")
	if err != nil {
		return
	}

	//read straightforward
	var headA, headB byte
	binary.Read(file, binary.LittleEndian, &headA)
	binary.Read(file, binary.LittleEndian, &headB)

	var size uint32
	binary.Read(file, binary.LittleEndian, &size)

	var reservedA, reservedB uint16
	binary.Read(file, binary.LittleEndian, &reservedA)
	binary.Read(file, binary.LittleEndian, &reservedB)

	var offbits uint32
	binary.Read(file, binary.LittleEndian, &offbits)

	fmt.Println(headA, headB, size, reservedA, reservedB, offbits)


	//use struct
	infoHeader := new(BitmapInfoHeader)
	binary.Read(file, binary.LittleEndian, infoHeader)

	fmt.Println(infoHeader)
}
