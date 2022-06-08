package handlers

func (u *DataHandler) SrartServer() error {
	u.Routers()
	return u.router.Run(":8080")
}
