package sourceservice

func (s *Service) DeleteSource(id string) error {
	if err := s.repo.DeleteSource(id); err != nil {
		return err
	}

	return nil
}
