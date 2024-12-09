package main

import (
	"aoc/libaoc"
	"fmt"
)

// Block is map[POS]fileid (empty = -1)
type Blocks map[int]int

// IdMaps is map[ID]pos,pos,pos
type IdMaps map[int][]int

type Disk struct {
	blocks        Blocks
	checksum      int
	emptySpaces   IdMaps
	fileLocations IdMaps
}

func (d *Disk) build(input []string) {
	d.blocks = make(Blocks)
	d.emptySpaces = make(IdMaps)
	d.fileLocations = make(IdMaps)

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
				d.fileLocations[fileID] = append(d.fileLocations[fileID], posWrite)
				posWrite++
			}
			fileID++
			databit = !databit
		case false:
			dataLength := libaoc.SilentAtoi(string(input[0][posInput]))
			emptyID := len(d.emptySpaces)
			for data := 0; data < dataLength; data++ {
				d.blocks[posWrite] = -1
				d.emptySpaces[emptyID] = append(d.emptySpaces[emptyID], posWrite)
				posWrite++
			}
			databit = !databit
		}
	}
}

// defragWholeFiles recursively moves files to the left most possible place. Start with -1!
func (d *Disk) defragWholeFiles(fileID int) {
	if fileID == -1 {
		fileID = len(d.fileLocations) - 1
	}
	for emptySpace := 0; emptySpace < len(d.emptySpaces); emptySpace++ {
		if len(d.emptySpaces[emptySpace]) >= len(d.fileLocations[fileID]) {
			// First we check if this would mean moving the file farther right!
			if d.emptySpaces[emptySpace][0] > d.fileLocations[fileID][0] {
				continue
			}
			// We write empty (-1) where the file used to be in the blocks
			previousPosition := d.fileLocations[fileID]
			for _, position := range previousPosition {
				d.blocks[position] = -1
			}
			// We write the file on the new location in the blocks and change the positions accordingly
			for fileBit := 0; fileBit < len(d.fileLocations[fileID]); fileBit++ {
				d.blocks[d.emptySpaces[emptySpace][0+fileBit]] = fileID
				d.fileLocations[fileID][fileBit] = d.blocks[d.emptySpaces[emptySpace][0+fileBit]]
			}
			// fmt.Printf("We've moved %d\n", fileID)
			d.recalcEmptySpaces()
			// d.print()
			if fileID > 0 {
				d.defragWholeFiles(fileID - 1)
			}
			return
		}
	}
	// appears we didn't find a big enough space
	// fmt.Printf("We couldn't move %d", fileID)
	if fileID > 0 {
		// fmt.Printf(", and continue with %d\n", fileID-1)
		d.defragWholeFiles(fileID - 1)
	}
	// fmt.Printf("\n")
	return
}

func (d *Disk) recalcEmptySpaces() {
	// fmt.Printf("Before: %#v\n", d.emptySpaces)
	d.emptySpaces = make(IdMaps)
	emptyID := 0
	writingEmpty := false
	for pos := 0; pos < len(d.blocks); pos++ {
		if d.blocks[pos] == -1 {
			writingEmpty = true
			d.emptySpaces[emptyID] = append(d.emptySpaces[emptyID], pos)
		} else {
			if writingEmpty {
				emptyID++
				writingEmpty = false
			}
		}
	}
	// fmt.Printf("After: %#v\n", d.emptySpaces)
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
	d.print()
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
			d.print()
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
