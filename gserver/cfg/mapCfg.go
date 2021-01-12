package cfg

type mapinfo struct {
	IndexCfg []int `json:"indexCfg"`
	Tilesize struct {
		Width  float64 `json:"width"`
		Height int     `json:"height"`
	} `json:"tilesize"`
	Movespeed int `json:"movespeed"`
	Metatype  int `json:"metatype"`
	Areas     []struct {
		SegInfo    []int  `json:"segInfo"`
		Setindex   int    `json:"setindex"`
		AreasIndex []int  `json:"areasIndex"`
		SetList    []int  `json:"setList"`
		Name       string `json:"name"`
		Beside     []int  `json:"beside"`
		Type       int    `json:"type"`
	} `json:"areas"`
	PointsAry [][]int `json:"pointsAry"`
	ArrowAry1 [][]int `json:"arrowAry1"`
	ArrowAry2 [][]int `json:"arrowAry2"`
	MaxX      int     `json:"maxX"`
	MaxY      int     `json:"maxY"`
	Mapsize   struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"mapsize"`
}
