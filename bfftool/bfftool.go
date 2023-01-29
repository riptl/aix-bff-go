package main

import (
	"bufio"
	"encoding/binary"
	"flag"
	"fmt"
	bff "github.com/terorie/aix-bff-go"
	"io"
	"os"
	"path/filepath"
)

var outDir string

func main() {
	flag.StringVar(&outDir, "o", "", "output directory")
	flag.Parse()

	if flag.NArg() != 1 {
		fmt.Fprintln(os.Stderr, "usage: bfftool <file>")
		os.Exit(1)
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()

	if err := dumpArchive(file); err != nil {
		panic(err.Error())
	}
}

func dumpArchive(file *os.File) error {
	bufrd := bufio.NewReader(file)

	var fileHdr bff.FileHeader
	if err := binary.Read(bufrd, binary.BigEndian, &fileHdr); err != nil {
		return fmt.Errorf("failed to read file header: %w", err)
	}
	if fileHdr.Magic != bff.Magic {
		return fmt.Errorf("not a BFF file")
	}

	for {
		var header bff.RecordHeader
		if err := binary.Read(bufrd, binary.LittleEndian, &header); err != nil {
			return fmt.Errorf("failed to read record beader: %w", err)
		}

		name, err := bff.ReadAlignedString(bufrd)
		if err != nil {
			return fmt.Errorf("failed to read record name: %w", err)
		}
		fmt.Println(name)

		var trailer bff.RecordTrailer
		if err := binary.Read(bufrd, binary.LittleEndian, &trailer); err != nil {
			return fmt.Errorf("failed to read record trailer: %w", err)
		}

		if header.Size > 0 {
			if err := dumpRecord(bufrd, name, header.Size); err != nil {
				return err
			}
		}

		alignedUp := (header.Size + 7) &^ 7
		if _, err := bufrd.Discard(int(alignedUp - header.Size)); err != nil {
			return fmt.Errorf("failed to skip record data padding: %w", err)
		}
	}

	return nil
}

func dumpRecord(rd io.Reader, name string, sz uint32) error {
	name = filepath.Join(outDir, name)
	dir := filepath.Dir(name)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("failed to create dirs: %w", err)
	}
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()
	_, err = io.Copy(file, io.LimitReader(rd, int64(sz)))
	return err
}
