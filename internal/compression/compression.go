package compression

import (
	"bytes"
	"compress/flate"
	"encoding/json"
	"fmt"
	"tile-game/internal/game"
	"unsafe"
)

func CompressUsingDictionary(h game.History, index int) (bytes.Buffer, error) {
	fmt.Println(h, "", index)

	var hdict game.History
	var hnew game.History

	if index == 0 {
		fmt.Println("branch == 0")
		var buff bytes.Buffer
		zw, _ := flate.NewWriter(&buff, flate.BestCompression)
		new, _ := json.Marshal(h)
		zw.Write(new)
		zw.Close()

		fmt.Printf("total JSON size: %d \n", unsafe.Sizeof(h))
		fmt.Printf("Compressed size: %d bytes\n", unsafe.Sizeof(buff))
		return buff, nil
	}

	fmt.Println("branch >= 1")

	hdict = game.History{
		States: h.States[:index-1],
		Moves:  h.Moves[:index-1],
	}

	hnew = game.History{
		States: h.States[index-1:],
		Moves:  h.Moves[index-1:],
	}

	fmt.Printf("hdict %d %d\n", len(hdict.States), len(hdict.Moves))
	fmt.Printf("hnew %d %d\n", len(hnew.States), len(hnew.Moves))

	//make(game.History(h.States[:index], ))
	dict, _ := json.Marshal(hdict)
	fmt.Println(hdict)

	fmt.Printf("#### \n\n")

	new, _ := json.Marshal(hnew)
	fmt.Println(hnew)

	var buff bytes.Buffer
	zw, _ := flate.NewWriterDict(&buff, flate.BestCompression, dict)
	zw.Write(new)
	zw.Close()

	fmt.Printf("total JSON size: (%d + %d)\n", len(dict), len(new))
	fmt.Printf("Compressed size: %d bytes\n", buff.Len())

	/*
		zr := flate.NewReaderDict(&buff, dict)

		var decompressedBuff bytes.Buffer
		nbytes, _ := io.Copy(&decompressedBuff, zr)

		fmt.Println(nbytes)
		fmt.Println(bytes.Equal(decompressedBuff.Bytes(), new))
	*/

	return buff, nil
}
