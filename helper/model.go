package helper

import (
	"github.com/nuraziz04/echo-restful-api-v2/model/domain"
	"github.com/nuraziz04/echo-restful-api-v2/model/web"
)

func ToPegawaiResponse(pg domain.Pegawai) web.PegawaiResponse {
	return web.PegawaiResponse{
		Id:      pg.Id,
		Name:    pg.Name,
		Alamat:  pg.Alamat,
		Telepon: pg.Telepon,
	}
}

func ToPegawaiResponses(pgs []domain.Pegawai) []web.PegawaiResponse {
	var pegawais []web.PegawaiResponse
	for _, pg := range pgs {
		pegawais = append(pegawais, ToPegawaiResponse(pg))
	}
	return pegawais
}
