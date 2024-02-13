package server

import (
	"log"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func NewServer(h *Handler) {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)

	r.GET("/projects/:id/amountPerDay", h.AmountPerDay)
	r.GET("/projects/:id/contributionsPerDay", h.ContributionsPerDay)
	r.GET("/projects/:id/contributorsPerDay", h.ContributorsPerDay)
	r.GET("/projects/:id/averageContribution", h.AverageContribution)
	r.GET("/projects/:id/percentagePerDay", h.PercentagePerDay)
	r.GET("/projects/:id/visitsPerDay", h.VisitsPerDay)
	r.GET("/projects/:id/conversionRatePerDay", h.ConversionRatePerDay)
	r.GET("/projects/:id/contributionMilestones/:milestone", h.ContributionMilestones)
	url := ginSwagger.URL("http://localhost:3000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	if err := r.Run(":3000"); err != nil {
		log.Fatal("error serving http", err)
	}
}
