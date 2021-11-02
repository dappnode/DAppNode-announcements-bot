package file

import (
	"encoding/json"
	"io/ioutil"
)

// define data structure
type Repository struct {
	Name string
	Address string
}
type Repositories struct {
	Repos []Repository
}

func GetRepos() ([]Repository, error) {
	// Open our jsonFile
    data, err := ioutil.ReadFile("repositories.json")
    // if we os.Open returns an error then handle it
    if err != nil {
        return nil, err
    }

	// json data
	var jsonFile Repositories

	// unmarshall it
	err = json.Unmarshal(data, &jsonFile)
	if err != nil {
		return nil, err
	}

	return jsonFile.Repos, nil
}