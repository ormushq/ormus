package sourceservice

func (s Service) DeleteSource(id, userID string) error {
	return s.repo.DeleteSource(id, userID)
}
