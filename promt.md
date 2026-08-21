bentar bentar bro gw kan login pakai data dibawah ini kan yah
{
    "user":"manu",
    "password":"123"
}
nah status itu succes nih..

nah sedangkan code kita itu ini:
func modelBindingJson(c *gin.Context) {
	var json Login
	if err := c.ShouldBind(&json); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if json.User != "manu" || json.Password != "123" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": "unauthorized",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "selamat anda login loh yah"})
}

lalu bagaimana bro kok bisa menerima request nama dan password? sedangkan tidak ada http seperti json newencoder json newdecoder buat menerima request json, tolong jelaskan misalnya bila ada salah kata dari saya mungkin anda bisa membenarkan juga, semoga anda mengerti yah