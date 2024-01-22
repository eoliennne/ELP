package main

func Decoupage(n, width, height int) [][]int {
	var list [][]int //ligne = tranche, list[] = [xinf, xsup, yinf, ysup]

	if width >= height {
		//découpage en colonnes

		bande := width/n - 1 //faire division euclidienne
		reste := width % n
		var min, max int = 0, bande

		for i := 0; i < n-1; i++ {
			var arr []int
			arr = append(arr, min, max, 0, height-1)
			list = append(list, arr)
			min, max = max+1, max+bande+1
			//print(list[i])
		}
		var arr []int
		arr = append(arr, min, max+reste)
		list = append(list, arr)
	}
	if height > width {
		//découpage en lignes
		bande := height/n - 1 //faire division euclidienne
		reste := height % n
		var min, max int = 0, bande

		for i := 0; i < n-1; i++ {
			var arr []int
			arr = append(arr, 0, width-1, min, max)
			list = append(list, arr)
			min, max = max+1, max+bande+1
			print(list[i])
		}
		var arr []int
		arr = append(arr, min, max+reste)
		list = append(list, arr)
	}

	return list

}
