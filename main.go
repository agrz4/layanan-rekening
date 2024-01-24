package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var saldo = make(map[int]float64)
var nomorRekening = 1000

func buatTabungan(c *gin.Context) {
	nomorRekening++
	saldo[nomorRekening] = 0
	pesan := fmt.Sprintf("Tabungan baru telah dibuat. Nomor rekening Anda adalah: %d", nomorRekening)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"pesan": pesan,
	})
}

func lihatSaldo(c *gin.Context) {
	nomorStr := c.PostForm("nomor")
	nomor, err := strconv.Atoi(nomorStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"pesan": fmt.Sprintf("Nomor rekening harus berupa angka. Error: %s", err.Error()),
		})
		return
	}

	if _, ok := saldo[nomor]; ok {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"pesan": fmt.Sprintf("Saldo untuk nomor rekening %d adalah: %.2f", nomor, saldo[nomor]),
		})
	} else {
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"pesan": "Nomor rekening tidak ditemukan.",
		})
	}
}

func setoran(c *gin.Context) {
	nomorStr := c.PostForm("nomor")
	nomor, err := strconv.Atoi(nomorStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"pesan": fmt.Sprintf("Nomor rekening harus berupa angka. Error: %s", err.Error()),
		})
		return
	}

	jumlahStr := c.PostForm("jumlah")
	jumlah, err := strconv.ParseFloat(jumlahStr, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"pesan": fmt.Sprintf("Jumlah setoran harus berupa angka. Error: %s", err.Error()),
		})
		return
	}

	if _, ok := saldo[nomor]; ok {
		saldo[nomor] += jumlah
		c.HTML(http.StatusOK, "index.html", gin.H{
			"pesan": fmt.Sprintf("Setoran berhasil. Saldo terbaru untuk nomor rekening %d adalah: %.2f", nomor, saldo[nomor]),
		})
	} else {
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"pesan": "Nomor rekening tidak ditemukan.",
		})
	}
}

func penarikan(c *gin.Context) {
	nomorStr := c.PostForm("nomor")
	nomor, err := strconv.Atoi(nomorStr)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"pesan": fmt.Sprintf("Nomor rekening harus berupa angka. Error: %s", err.Error()),
		})
		return
	}

	jumlahStr := c.PostForm("jumlah")
	jumlah, err := strconv.ParseFloat(jumlahStr, 64)
	if err != nil {
		c.HTML(http.StatusBadRequest, "index.html", gin.H{
			"pesan": fmt.Sprintf("Jumlah penarikan harus berupa angka. Error: %s", err.Error()),
		})
		return
	}

	if _, ok := saldo[nomor]; ok {
		if saldo[nomor] >= jumlah {
			saldo[nomor] -= jumlah
			c.HTML(http.StatusOK, "index.html", gin.H{
				"pesan": fmt.Sprintf("Penarikan berhasil. Saldo terbaru untuk nomor rekening %d adalah: %.2f", nomor, saldo[nomor]),
			})
		} else {
			c.HTML(http.StatusBadRequest, "index.html", gin.H{
				"pesan": "Saldo tidak mencukupi.",
			})
		}
	} else {
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"pesan": "Nomor rekening tidak ditemukan.",
		})
	}
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"pesan": "Selamat datang! Silakan pilih layanan.",
		})
	})

	router.POST("/buat-tabungan", buatTabungan)
	router.POST("/lihat-saldo", lihatSaldo)
	router.POST("/setoran", setoran)
	router.POST("/penarikan", penarikan)

	router.Run(":8080")
}
