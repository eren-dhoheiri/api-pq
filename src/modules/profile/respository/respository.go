package respository

import (
	"github.com/eren-dhoheiri/api-pq/src/modules/profile/model"
)

//ProfileRespository
type ProfileRespository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindAll() (*model.Profiles, error)
}
