package main

func findJudge(n int, trust [][]int) int {
	if n == 1 {
		return 1
	}
	m := map[int]int{}

	judge := 0
	max := 0

	for i := 0; i < len(trust); i++ {
		if judge == trust[i][0] {
			judge = -1
			max = 0
		}
		if m[trust[i][1]] == -1 {
			continue
		}

		m[trust[i][1]] += 1

		if m[trust[i][1]] > max {
			judge = trust[i][1]
			max = m[trust[i][1]]
		}
	}

	if max != n {
		return -1
	}

	return judge
}
