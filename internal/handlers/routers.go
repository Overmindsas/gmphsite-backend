package handlers

func (u *DataHandler) Routers() {
	u.router.GET("/getdataresp", u.GetData)
	u.router.GET("/getalldata", u.GetAllData)
}
