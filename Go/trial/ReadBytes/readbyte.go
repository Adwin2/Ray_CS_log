package main

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func CreateAndWrite(filename string) {
	f, err := os.Create(filename)
	if  err != nil  {
		panic(err)
	}
	defer f.Close()

	binary.Write(f, binary.LittleEndian, []byte{
		0xFD,                       /* Indicates that this key ("baz") has an expire,
                            		   and that the expire timestamp is expressed in seconds. */
		0x52, 0xED, 0x2A, 0x66,              /* The expire timestamp, expressed in Unix time,
												stored as an 4-byte unsigned integer, in little-endian (read right-to-left).
												Here, the expire timestamp is 1714089298. */
		0x00,                       // Value type is string.
		0x03, 0x62, 0x61, 0x7A,              // Key name is "baz".
		0x03, 0x71, 0x75, 0x78,              // Value is "qux".
		0x15, 0x72, 0xE7, 0x07, 0x8F, 0x01, 0x00, 0x00,
	})
}

func Read(filename string) {
	var res []string
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := bufio.NewReader(f)
	n, err := io.CopyN(io.Discard, reader, 5)
	if err != nil {
		panic(err)
	}
	if n < 5  {
		log.Printf("不够长")
	}

	for{
		_, err := reader.ReadByte()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		keylen, err := reader.ReadByte()
		if err == io.EOF {
			break
		}	
		if err != nil {
			panic(err)
		}
		// fmt.Println(flag, keylen)
		keyBytes := make([]byte, keylen)
		_, err = io.ReadFull(reader, keyBytes)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			break
		}
		if err != nil {
			panic(err)
		}
		res = append(res, string(keyBytes))

		vallen, err := reader.ReadByte()
		valBytes := make([]byte, vallen)
		_, err = io.ReadFull(reader, valBytes)
		if err == io.EOF || err == io.ErrUnexpectedEOF  {
			break
		}
		if err != nil {
			panic(err)
		}
		res = append(res, string(valBytes))
		buf := make([]byte, 8)
		_, err = io.ReadFull(reader, buf)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return
		}
		if err != nil {
			panic(err)
		}
		ts := binary.LittleEndian.Uint64(buf)
		fmt.Println(ts)
	}
	fmt.Println(res)
}

func main() {
	filename := "/home/raymond/a.rdb"
	CreateAndWrite(filename)
	Read(filename)

	fmt.Println(filepath.Dir(filename))
	// for range 10 {
	// 	fmt.Println(1)
	// }
}