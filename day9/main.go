package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	//partOne()
	partTwo()
}

func partOne() {
	diskMap := readLines("input.txt")
	println("Diskmap: ", diskMap)

	compactDiskMap := compactFiles(diskMap)
	println("Compact Diskmap: ", compactDiskMap)

	defraggedDiskMap := rearrangeFileBlocks(compactDiskMap)
	println("Defragged diskmap: ", defraggedDiskMap)

	checksum := calculateChecksum(defraggedDiskMap)
	println("Checksum: ", checksum)
}

func partTwo() {
	diskMap := readLines("input.txt")
	println("Diskmap: ", diskMap)

	compactDiskMap := compactFiles(diskMap)
	println("Compact Diskmap: ", compactDiskMap)

	defraggedDiskMap := rearrangeFiles(compactDiskMap)
	println("Defragged diskmap: ", defraggedDiskMap)

	checksum := calculateChecksum(defraggedDiskMap)
	println("Checksum: ", checksum)
}

func calculateChecksum(diskMap string) int {
	splittedDiskMap := strings.Split(diskMap, ",")

	checksum := 0
	for i, block := range splittedDiskMap {
		if block != "." && block != "," {
			fileID, _ := strconv.Atoi(block)
			checksum += i * fileID
		}
	}
	return checksum
}

func rearrangeFileBlocks(diskMap string) string {
	defraggedDiskMap := strings.Split(diskMap, ",")
	reverseIndex := len(defraggedDiskMap) - 1

	for i := 0; i < len(defraggedDiskMap); i++ {
		if string(defraggedDiskMap[i]) == "." {
			for j := reverseIndex; j > i; j-- {
				if isNumeric(string(defraggedDiskMap[j])) {
					defraggedDiskMap[i] = defraggedDiskMap[j]
					defraggedDiskMap[j] = "."
					break
				}
			}
		}
	}

	return strings.Join(defraggedDiskMap, ",")
}

func findFirstMatchingDotsIndex(diskMap []string, length int) int {
	for i := 0; i < len(diskMap); i++ {
		dotLength := 0
		for k := i; k < len(diskMap) && diskMap[k] == "."; k++ {
			dotLength++
		}

		if dotLength >= length {
			return i
		}
	}
	return -1
}

func rearrangeFiles(diskMap string) string {
	defraggedDiskMap := strings.Split(diskMap, ",")
	reverseIndex := len(defraggedDiskMap) - 1

	for j := reverseIndex; j > 0; j-- {
		if isNumeric(defraggedDiskMap[j]) {
			var file []string
			file = append(file, defraggedDiskMap[j])

			//get the complete file
			for fileComplete := false; !fileComplete; {
				if j == 0 {
					break
				}

				if defraggedDiskMap[j-1] == defraggedDiskMap[j] || defraggedDiskMap[j-1] == "," {
					file = append(file, defraggedDiskMap[j-1])
					j--
					continue
				}

				fileComplete = true
			}

			dotsIndex := findFirstMatchingDotsIndex(defraggedDiskMap, len(file))

			//not enough room to replace file
			if dotsIndex == -1 || dotsIndex > j {
				continue
			}

			for i := dotsIndex; i < dotsIndex+len(file); i++ {
				defraggedDiskMap[i] = file[0]
			}

			for i := j; i < j+len(file); i++ {
				defraggedDiskMap[i] = "."
			}
		}
	}

	return strings.Join(defraggedDiskMap, ",")
}

func isNumeric(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func compactFiles(diskMap string) string {
	blocks := strings.Split(diskMap, "")
	numberOfBlocks := len(blocks)
	var result []string

	for i := 0; i < numberOfBlocks; i += 2 {
		fileLength, _ := strconv.Atoi(string(blocks[i]))
		freeSpaceLength := 0

		if i+1 < numberOfBlocks {
			freeSpaceLength, _ = strconv.Atoi(string(blocks[i+1]))
		}

		for j := 0; j < fileLength; j++ {
			result = append(result, strconv.Itoa(i/2))
		}

		for j := 0; j < freeSpaceLength; j++ {
			result = append(result, ".")
		}
	}

	return strings.Join(result, ",")
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) string {
	content, err := os.ReadFile(path)
	text := string(content)
	if err != nil {
		log.Fatal(err)
	}

	return text
}
