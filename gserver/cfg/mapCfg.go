package cfg

type mapinfo struct {
	Areas []struct {
		Setindex   int    `json:"setindex"`
		AreasIndex []int  `json:"areasIndex"`
		SegInfo    []int  `json:"segInfo"`
		Name       string `json:"name"`
		Type       int    `json:"type"`
	} `json:"areas"`
	PointsAry [][]float64 `json:"pointsAry"`
	ArrowAry1 [][]int     `json:"arrowAry1"`
	ArrowAry2 [][]int     `json:"arrowAry2"`
	MaxX      int         `json:"maxX"`
	MaxY      int         `json:"maxY"`
	Mapsize   struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	} `json:"mapsize"`
	Tilesize struct {
		Width  int     `json:"width"`
		Height float64 `json:"height"`
	} `json:"tilesize"`
	Movespeed int `json:"movespeed"`
	Metatype  int `json:"metatype"`
}
