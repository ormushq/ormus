package sourcevalidator_test

import (
	"fmt"
	"testing"

	"github.com/ormushq/ormus/manager/entity"
	"github.com/ormushq/ormus/manager/mock/sourcemock"
	"github.com/ormushq/ormus/manager/param"
	"github.com/ormushq/ormus/manager/validator/sourcevalidator"
	"github.com/stretchr/testify/assert"
)

func TestValidateUpdateSourceForm(t *testing.T) {
	testCases := []struct {
		name    string
		params  param.UpdateSourceRequest
		repoErr bool
		error   error
	}{
		{
			name:  "less than min name len",
			error: fmt.Errorf("Name: the length must be between 5 and 30\n"),
			params: param.UpdateSourceRequest{
				Name:        "le",
				Description: "new normal description",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
				Status:      entity.StatusNotActive,
			},
		},
		{
			name:  "more than max name len",
			error: fmt.Errorf("Name: the length must be between 5 and 30\n"),
			params: param.UpdateSourceRequest{
				Name:        "more than max name len la la la la la la la la la la la la la la la la",
				Description: "new normal description",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
				Status:      entity.StatusNotActive,
			},
		},
		{
			name:  "less than min description len",
			error: fmt.Errorf("Description: the length must be between 5 and 100\n"),
			params: param.UpdateSourceRequest{
				Name:        "normal new name",
				Description: "de",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
				Status:      entity.StatusNotActive,
			},
		},
		{
			name:  "more than max description len",
			error: fmt.Errorf("Description: the length must be between 5 and 100\n"),
			params: param.UpdateSourceRequest{
				Name:        "normal new name",
				Description: "more then max description len la la la la la la la la la la la la la lal la la lal al al lal ala lal al lala l l",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
				Status:      entity.StatusNotActive,
			},
		},
		{
			name:  "invalide project id",
			error: fmt.Errorf("ProjectID: invalid id\n"),
			params: param.UpdateSourceRequest{
				Name:        "normal new name",
				Description: "new normal description",
				ProjectID:   "invalide project id",
				Status:      entity.StatusNotActive,
			},
		},
		{
			name:  "ordinary",
			error: nil,
			params: param.UpdateSourceRequest{
				Name:        "normal new name",
				Description: "new normal description",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
				Status:      entity.StatusNotActive,
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			mockRepo := sourcemock.NewMockRepository(tc.repoErr)
			valid := sourcevalidator.New(mockRepo)

			// 2. execution
			res := valid.ValidateUpdateSourceForm(tc.params)

			// 3. assertion
			if tc.error == nil {
				assert.Nil(t, res)
				return
			}
			assert.Equal(t, tc.error.Error(), res.Error())
		})
	}
}

func TestValidateCreateSourceForm(t *testing.T) {
	defaulte := sourcemock.DefaultSource()

	testCases := []struct {
		name    string
		params  param.AddSourceRequest
		repoErr bool
		error   error
	}{
		{
			name:  "less than min name len",
			error: fmt.Errorf("Name: the length must be between 5 and 30\n"),
			params: param.AddSourceRequest{
				Name:        "le",
				Description: "new normal description",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
			},
		},
		{
			name:  "more than max name len",
			error: fmt.Errorf("Name: the length must be between 5 and 30\n"),
			params: param.AddSourceRequest{
				Name:        "more than max name len la la la la la la la la la la la la la la la la",
				Description: "new normal description",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
			},
		},
		{
			name:  "less than min description len",
			error: fmt.Errorf("Description: the length must be between 5 and 100\n"),
			params: param.AddSourceRequest{
				Name:        "normal new name",
				Description: "de",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
			},
		},
		{
			name:  "more than max description len",
			error: fmt.Errorf("Description: the length must be between 5 and 100\n"),
			params: param.AddSourceRequest{
				Name:        "normal new name",
				Description: "more then max description len la la la la la la la la la la la la la lal la la lal al al lal ala lal al lala l l",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
			},
		},
		{
			name:  "invalide project id",
			error: fmt.Errorf("ProjectID: invalid id\n"),
			params: param.AddSourceRequest{
				Name:        "normal new name",
				Description: "new normal description",
				ProjectID:   "invalide project id",
			},
		},
		{
			name:  "exist source",
			error: fmt.Errorf("Name: this name is already usesd\n"),
			params: param.AddSourceRequest{
				Name:        defaulte.Name,
				Description: "new normal description",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
			},
		},
		{
			name:  "ordinary",
			error: nil,
			params: param.AddSourceRequest{
				Name:        "normal new name",
				Description: "new normal description",
				ProjectID:   "01HJDQ386MW8EM6WMC8B6J5HAN",
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 1. setup
			mockRepo := sourcemock.NewMockRepository(tc.repoErr)
			valid := sourcevalidator.New(mockRepo)

			// 2. execution
			res := valid.ValidateCreateSourceForm(tc.params)

			// 3. assertion
			if tc.error == nil {
				assert.Nil(t, res)
				return
			}
			assert.Equal(t, tc.error.Error(), res.Error())
		})
	}
}
