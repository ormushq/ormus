package sourceservice

func (s Service) DeleteSource(id string) error {
	return s.repo.DeleteSource(id)
}
