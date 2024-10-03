package main

func calculateNextState(p golParams, world [][]byte) [][]byte {

	sum := 0

	imgHt := p.imageHeight
	imgWd := p.imageWidth

	newWorld := make([][]byte, imgHt)
	for i, _ := range newWorld {
		newWorld[i] = make([]byte, imgWd)
	}

	for y := 0; y < imgHt; y++ {
		for x := 0; x < imgWd; x++ {
			topLeft := int(world[(y-1+imgHt)%imgHt][(x-1+imgWd)%imgWd]) % 254
			topMiddle := int(world[(y-1+imgHt)%imgHt][(x+imgWd)%imgWd]) % 254
			topRight := int(world[(y-1+imgHt)%imgHt][(x+1+imgWd)%imgWd]) % 254

			middleLeft := int(world[(y+imgHt)%imgHt][(x-1+imgWd)%imgWd]) % 254
			middleRight := int(world[(y+imgHt)%imgHt][(x+1+imgWd)%imgWd]) % 254

			bottomLeft := int(world[(y+1+imgHt)%imgHt][(x-1+imgWd)%imgWd]) % 254
			bottomMiddle := int(world[(y+1+imgHt)%imgHt][(x+imgWd)%imgWd]) % 254
			bottomRight := int(world[(y+1+imgHt)%imgHt][(x+1+imgWd)%imgWd]) % 254

			sum =
				topLeft + topMiddle + topRight +
					middleLeft + middleRight +
					bottomLeft + bottomMiddle + bottomRight

			if int(world[y][x]) > 0 {
				if sum < 2 {
					newWorld[y][x] = 0
				} else if sum == 2 || sum == 3 {
					newWorld[y][x] = 255
				} else {
					newWorld[y][x] = 0

				}
			} else {

				if sum == 3 {
					newWorld[y][x] = 255

				} else {
					newWorld[y][x] = 0
				}
			}

		}
	}

	return newWorld
}

func calculateAliveCells(p golParams, world [][]byte) []cell {

	aliveCells := []cell{}

	for y := 0; y < p.imageHeight; y++ {
		for x := 0; x < p.imageWidth; x++ {
			if world[y][x] > 0 {
				aliveCells = append(aliveCells, cell{x: x, y: y})
			}
		}
	}

	return aliveCells
}
