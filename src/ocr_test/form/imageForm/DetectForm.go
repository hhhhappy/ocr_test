package imageForm

type DetectForm struct {
	ImageUrl    string `form:"image_url" json:"image_url"`
	ImageBase64 string `form:"image_base64" json:"image_base64"`
}
