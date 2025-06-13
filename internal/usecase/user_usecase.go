package usecase

type UserUsecase struct {
	ImageProcessingUsecase ImageProcessingUsecase
}

func NewUserUsecase(usecase ImageProcessingUsecase) *UserUsecase {
	return &UserUsecase{
		ImageProcessingUsecase: usecase,
	}
}

func (c *ImageProcessingUsecase) UpdateProfilePicture() {

}
