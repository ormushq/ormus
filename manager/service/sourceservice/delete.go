package sourceservice

func (s Service) DeleteSource(id string, userID string) error {
	return s.repo.DeleteSource(id, userID)
}
