package handler

// if user.Username == "" {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username is required"))
// 	return
// } else if user.Password == "" {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password is required"))
// 	return
// } else if user.FullName == "" {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Full name is required"))
// 	return
// } else if user.Email == "" {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is required"))
// 	return
// }

// if !h.service.CheckPassLength(user.Password) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must be at least 8 characters"))
// 	return
// } else if !h.service.HasUpperLetter(user.Password) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 uppercase letter"))
// 	return
// } else if !h.service.HasLowerLetter(user.Password) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 lowercase letter"))
// 	return
// } else if !h.service.HasNumber(user.Password) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 number"))
// 	return
// } else if !h.service.HasSpecialChar(user.Password) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Password must contain at least 1 special character"))
// 	return
// }

// if !h.service.CheckEmail(user.Email) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Email is not valid"))
// 	return
// } else if !h.service.CheckUsername(user.Username) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Username should not contain spaces"))
// 	return
// } else if !h.service.CheckFullName(user.FullName) {
// 	c.JSON(http.StatusBadRequest, model.NewErrorResponse(http.StatusBadRequest, "Full name should not contain numbers or special characters"))
// 	return
// }