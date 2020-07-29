package main

import (
	"fmt"

	"github.com/eren-dhoheiri/api-pq/config"
	"github.com/eren-dhoheiri/api-pq/src/modules/profile/model"
	"github.com/eren-dhoheiri/api-pq/src/modules/profile/respository"
)

func main() {
	fmt.Println("Hello")

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	err = db.Ping()

	if err != nil {
		fmt.Println(err)
	}

	profileRespositoryPostgres := respository.NewProfileRespositoryPostgres(db)

	profiles, err := getProfiles(profileRespositoryPostgres)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range profiles {
		fmt.Println(v)
	}

}

func saveProfile(p *model.Profile, repo respository.ProfileRespository) error {
	err := repo.Save(p)

	if err != nil {
		return err
	}

	return nil
}

func updateProfile(p *model.Profile, repo respository.ProfileRespository) error {
	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}

	return nil
}

func deleteProfile(id string, repo respository.ProfileRespository) error {
	err := repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func getProfile(id string, repo respository.ProfileRespository) (*model.Profile, error) {
	profile, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func getProfiles(repo respository.ProfileRespository) (model.Profiles, error) {
	profiles, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
