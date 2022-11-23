package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/theojulienne/go-wireless"
)

func notImplemented(c *gin.Context) {
	c.AbortWithStatus(501)
}

func json(err error) map[string]string {
	return map[string]string{"error": err.Error()}
}

func errStatus(err error) int {
	switch err {
	case wireless.ErrFailBusy:
		return 409
	default:
		return 500
	}
}

func listInterfaces(c *gin.Context) {
	c.JSON(200, wireless.Interfaces())
}

func listAccesPoints(c *gin.Context) {
	iface := c.Param("iface")

	wc, err := wireless.NewClient(iface)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(errStatus(err), json(err))
		return
	}
	defer wc.Close()

	wc.ScanTimeout = 3 * time.Second

	aps, err := wc.Scan()
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(errStatus(err), json(err))
		return
	}

	c.JSON(200, aps)
}

func listNetworks(c *gin.Context) {
	iface := c.Param("iface")

	wc, err := wireless.NewClient(iface)
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(errStatus(err), json(err))
		return
	}
	defer wc.Close()

	nets, err := wc.Networks()
	if err != nil {
		c.Error(err)
		c.AbortWithStatusJSON(errStatus(err), json(err))
		return
	}

	c.JSON(200, nets)
}
