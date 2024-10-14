package whiteboardws

type linedata struct {
	Points      []float64 `json:"points" binding:"required"`
	Stroke      string    `json:"stroke" binding:"required"`
	StrokeWidth int       `json:"strokeWidth" binding:"required"`
	Tool        string    `json:"tool" binding:"required"`
	Text        string    `json:"text"`
}
