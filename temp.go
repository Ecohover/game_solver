type YCoord int

const (
    A YCoord = iota
    B
    C
    D
    E
    F
)

type Cell struct {
    position [2]int // position[0] 為 X 座標，position[1] 為 Y 座標
    filled bool
}

type Board struct {
    size int
    cells [][]*Cell
}

// 定義一個 6*6 的棋盤
var board = Board{
    size: 6,
    cells: [][]*Cell{},
}

// 初始化棋盤
for i := 0; i < board.size; i++ {
    row := []*Cell{}
    for j := 0; j < board.size; j++ {
        row = append(row, &Cell{position: [i, YCoord(j)], filled: false})
    }
    board.cells = append(board.cells, row)
}
