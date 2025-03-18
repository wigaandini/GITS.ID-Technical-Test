package main

import (
	"fmt"
)

func hitungRankGITS(leaderboard []int, skorGits []int) []int {
	// biar gada dup skor
	temp := map[int]bool{}
	skorUnik := []int{}
	
	for i := 0; i < len(leaderboard); i++ {
		nilai := leaderboard[i]
		if temp[nilai] == false {
			temp[nilai] = true
			skorUnik = append(skorUnik, nilai)
		}
	}
	
	// sort skor dari besar ke kecil
	for i := 0; i < len(skorUnik)-1; i++ {
		for j := 0; j < len(skorUnik)-i-1; j++ {
			if skorUnik[j] < skorUnik[j+1] {
				temp := skorUnik[j]
				skorUnik[j] = skorUnik[j+1]
				skorUnik[j+1] = temp
			}
		}
	}
	
	// map skor ke rank
	rankNilai := map[int]int{}
	for i := 0; i < len(skorUnik); i++ {
		rankNilai[skorUnik[i]] = i + 1
		// fmt.Println(rankNilai)
	}

	// fmt.Println(skorUnik)
	// fmt.Println(rankNilai)
	
	hasilRank := make([]int, len(skorGits))
	
	for i := 0; i < len(skorGits); i++ {
		skor := skorGits[i]
		rank, ada := rankNilai[skor]
		
		if ada == true {
			hasilRank[i] = rank
		} else {
			if len(skorUnik) == 0 || skor > skorUnik[0] {
				hasilRank[i] = 1
			} else if skor < skorUnik[len(skorUnik)-1] {
				hasilRank[i] = len(skorUnik) + 1
			} else {
				for j := 0; j < len(skorUnik)-1; j++ {
					if skor < skorUnik[j] && skor >= skorUnik[j+1] {
						hasilRank[i] = j + 2
						break
					}
				}
			}
		}
	}
	
	return hasilRank
}

func main() {
	var n int
	fmt.Scan(&n)
	
	leaderboard := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&leaderboard[i])
	}
	
	var m int
	fmt.Scan(&m)
	
	skorGits := make([]int, m)
	for i := 0; i < m; i++ {
		fmt.Scan(&skorGits[i])
	}
	
	hasil := hitungRankGITS(leaderboard, skorGits)
	
	for i := 0; i < len(hasil); i++ {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(hasil[i])
	}
	fmt.Println()
}