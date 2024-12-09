package aoc2024

import (
	"advent/loader"
	"fmt"
	"strconv"
	"strings"
)

type FileWithSpace struct {
	FileId     int
	Size       int
	SpaceAfter int

	Start       int
	SpaceOffset int // After moving an item into the space after, this is the offset to the next unoccupied space
}

func (f *FileWithSpace) Checksum() int {
	return f.FileId * calculateRangeSum(f.Start, f.Start+f.Size-1)
}

func parseFilesWithSpace(lines []string) (*[]*FileWithSpace, error) {
	files := &[]*FileWithSpace{}
	joined := strings.Join(lines, "")
	for i := 0; i < len(joined); i += 2 {
		size, err := strconv.Atoi(string(joined[i]))
		if err != nil {
			return nil, err
		}
		space := 0
		if i+1 < len(joined) {
			space, err = strconv.Atoi(string(joined[i+1]))
			if err != nil {
				return nil, err
			}
		}

		file := &FileWithSpace{
			FileId:     i / 2,
			Size:       size,
			SpaceAfter: space,
		}
		*files = append(*files, file)
	}

	return files, nil
}

func calculateRangeSum(from, to int) int {
	return (to*(to+1) - (from-1)*from) / 2
}

func calculateChecksum(files *[]*FileWithSpace) int {
	checksum := 0
	position := 0
	current := 0
	last := len(*files) - 1

	for current <= last {
		currentItem := (*files)[current]
		if currentItem.Size == 0 {
			break
		}

		/*
			for x := 0; x < currentItem.Size; x++ {
				fmt.Printf("%d", currentItem.FileId)
			}
		*/

		// Calculate the current item's checksum contribution
		checksum += currentItem.FileId * calculateRangeSum(position, position+currentItem.Size-1)
		position += currentItem.Size

		if current == last {
			break
		}

		// move in enough items from the end to fill the space
		for currentItem.SpaceAfter > 0 {
			lastItem := (*files)[last]
			if lastItem.Size == 0 {
				last--
				continue
			}

			if lastItem.Size <= currentItem.SpaceAfter {
				/*
					for x := 0; x < lastItem.Size; x++ {
						fmt.Printf("%d", lastItem.FileId)
					}
				*/
				currentItem.SpaceAfter -= lastItem.Size
				checksum += lastItem.FileId * calculateRangeSum(position, position+lastItem.Size-1)
				position += lastItem.Size
				lastItem.Size = 0
				last--
			} else {
				/*
					for x := 0; x < currentItem.SpaceAfter; x++ {
						fmt.Printf("%d", lastItem.FileId)
					}
				*/
				lastItem.Size -= currentItem.SpaceAfter
				checksum += lastItem.FileId * calculateRangeSum(position, position+currentItem.SpaceAfter-1)
				position += currentItem.SpaceAfter
				currentItem.SpaceAfter = 0
			}
		}

		current++
	}
	fmt.Println()

	return checksum
}

func Day9Part1() {
	loader, err := loader.NewLoader("2024/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//loader.Lines = []string{"2333133121414131402"}

	items, err := parseFilesWithSpace(loader.Lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	checksum := calculateChecksum(items)

	fmt.Printf("Day 9 Part 1: %d\n", checksum)
}

func Day9Part2() {
	loader, err := loader.NewLoader("2024/day9.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	//loader.Lines = []string{"2333133121414131402"}

	items, err := parseFilesWithSpace(loader.Lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	currentEnd := 0
	for _, item := range *items {
		item.Start = currentEnd
		currentEnd += item.Size + item.SpaceAfter
	}

	// Working backwards, figure out if we have space for the file
	for i := len(*items) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if (*items)[j].SpaceAfter >= (*items)[i].Size {
				// Move the item
				(*items)[j].SpaceAfter -= (*items)[i].Size
				(*items)[i].Start = (*items)[j].Start + (*items)[j].Size + (*items)[j].SpaceOffset
				(*items)[j].SpaceOffset += (*items)[i].Size
				break
			}
		}
	}

	checksum := 0
	for _, item := range *items {
		checksum += item.Checksum()
	}

	fmt.Printf("Day 9 Part 2: %d\n", checksum)
}
