package main

import (
	"aoc/libaoc"
	"fmt"
)

// Block is map[POS]fileid (empty = -1)
type Blocks map[int]int

type Disk struct {
	blocks   Blocks
	checksum int
}

func (d *Disk) build(input []string) {
	d.blocks = make(Blocks)

	// Parse: Filelength, free space, file lenght, free space, ...
	databit := true
	posWrite := 0
	fileID := 0
	for posInput := 0; posInput < len(input[0]); posInput++ {
		switch databit {
		case true:
			dataLength := libaoc.SilentAtoi(string(input[0][posInput]))
			for data := 0; data < dataLength; data++ {
				d.blocks[posWrite] = fileID
				posWrite++
			}
			fileID++
			databit = !databit
		case false:
			dataLength := libaoc.SilentAtoi(string(input[0][posInput]))
			for data := 0; data < dataLength; data++ {
				d.blocks[posWrite] = -1
				posWrite++
			}
			databit = !databit
		}
	}
}

func (d *Disk) defrag() {
	seekPos := len(d.blocks) - 1
	// Is dataPos representing data? We need to have data in order to move it :O
	// I can't know of they would be so horrible to end with empty space
	for {
		if d.blocks[seekPos] != -1 {
			break
		} else {
			seekPos--
		}
	}
	// d.print()
	for defragPos := 0; defragPos < seekPos; defragPos++ {
		switch d.blocks[defragPos] {
		case -1:
			d.blocks[defragPos] = d.blocks[seekPos]
			d.blocks[seekPos] = -1
			seekPos--
			for {
				if d.blocks[seekPos] != -1 {
					break
				} else {
					seekPos--
				}
			}
			// d.print()
		default:
		}
	}
}

func (d *Disk) calcChecksum() {
	var sum int
	for pos := 0; pos < len(d.blocks); pos++ {
		switch d.blocks[pos] {
		case -1:
		default:
			sum += pos * d.blocks[pos]
		}
	}
	d.checksum = sum
	fmt.Printf("Checksum calculated: %d\n", d.checksum)
}

func (d *Disk) print() {
	for pos := 0; pos < len(d.blocks); pos++ {
		switch d.blocks[pos] {
		case -1:
			fmt.Printf(".")
		default:
			fmt.Printf("%d", d.blocks[pos])
		}
	}
	fmt.Printf("\n")
}
