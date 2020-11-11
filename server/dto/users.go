package dto

type GetProfileResp struct {
	ID              string `json:"id" example:"Yh34te-saaiud3322chadsc-asdvcsf"`
	Name            string `json:"name" example:"Tester"`
	Email           string `json:"email" example:"test@example.com"`
	AccountType     string `json:"account_type" example:"operator"`
	ProfileImageUrl string `json:"profile_image_url" example:"http://simple-service-backend/simple-service/photo-924y82hde7ce.jpg"`
} //@name GetProfileResp

type EditProfileReq struct {
	Name  string `json:"name" example:"Tester"`
	Email string `json:"email" example:"new-email@example.com"`
} //@name EditProfileReq
