package cmd

type S3Input struct {
	Item    []byte `json:"item"`
	ItemKey string `json:"itemKey"`
}
