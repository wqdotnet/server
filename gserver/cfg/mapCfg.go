package cfg

type mapinfo struct {
	Mapsize struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"mapsize"`
	Tilesize struct {
		Width  int     `json:"width"`
		Height float64 `json:"height"`
	} `json:"tilesize"`
	Movespeed int `json:"movespeed"`
	Metatype  int `json:"metatype"`
	Areas     []struct {
		Setindex   int    `json:"setindex"`
		Type       int    `json:"type"`
		SegInfo    []int  `json:"segInfo"`
		AreasIndex []int  `json:"areasIndex"`
		Beside     []int  `json:"beside"`
		Name       string `json:"name"`
	} `json:"areas"`
	PointsAry [][]float64 `json:"pointsAry"`
	ArrowAry1 [][]int     `json:"arrowAry1"`
	ArrowAry2 [][]int     `json:"arrowAry2"`
	MaxX      int         `json:"maxX"`
	MaxY      int         `json:"maxY"`
	IndexCfg  []int       `json:"indexCfg"`
}
