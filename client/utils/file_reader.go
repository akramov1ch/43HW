package utils

import (
    "encoding/json"
    "os"
    "client/models"
)

func ReadUsersFromFile(filename string) ([]models.User, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }

    var users []models.User
    err = json.Unmarshal(data, &users)
    if err != nil {
        return nil, err
    }

    return users, nil
}
