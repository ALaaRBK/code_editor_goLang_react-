package schemas

//Submission
type Submission struct {
	Lang      string `json:"lang"`
	CreatedAt int64  `json:"createdAt"`
	FilePath  string `json:"filePath"`
}
