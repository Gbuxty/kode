package handlers

import (
    "encoding/json"
    "net/http"
    "kode/models"
    "kode/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return
    }
    // Fetch user from database (pseudo code)
    // dbUser := db.FindUserByUsername(user.Username)
    dbUser := &models.User{ID: 1, Username: user.Username, Password: "$2a$14$Lc.UJ.UJ.UJ.UJ.UJ.UJ.UJ"} // Example hash
    err = dbUser.CheckPassword(user.Password)
    if err != nil {
        w.WriteHeader(http.StatusUnauthorized)
        return
    }
    token, err := utils.GenerateToken(dbUser.ID)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        return
    }
    w.Write([]byte(token))
}