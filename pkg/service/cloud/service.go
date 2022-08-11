package cloud

type Service interface {
	HostImageToGCS(imageUrl string, name string) (string, error)
}
