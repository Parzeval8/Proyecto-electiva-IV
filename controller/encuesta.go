package controller

import (
	"encoding/csv"
	"net/http"
	"os"
	"strconv"

	"strings"

	"github.com/Parzeval8/proyecto-api-go/config"
	"github.com/Parzeval8/proyecto-api-go/models"
	"github.com/gin-gonic/gin"
)

func Read_csv() {
	file, err := os.Open("csv/Mineria-numericos.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()

		var elements []string
		for _, str := range record {
			parts := strings.Split(str, ";")
			elements = append(elements, parts...)
		}
		record = elements
		if err != nil {
			break
		}

		iD, err := strconv.Atoi(record[0])
		uiD := uint(iD)
		p1Carrera, err := strconv.Atoi(record[1])
		p2Semestre, err := strconv.Atoi(record[2])
		p3Edad, err := strconv.Atoi(record[3])
		p4Sisben, err := strconv.Atoi(record[4])
		p5Transporte, err := strconv.Atoi(record[5])
		p6Orientacion, err := strconv.Atoi(record[6])
		p8Ideastu, err := strconv.Atoi(record[7])
		p9Perfil, err := strconv.Atoi(record[8])
		p12CambioU, err := strconv.Atoi(record[9])
		p15Trabajo, err := strconv.Atoi(record[10])
		p16Ingresos, err := strconv.Atoi(record[11])
		p17Reside, err := strconv.Atoi(record[12])
		p18Foraneo, err := strconv.Atoi(record[13])
		p19Pcargo, err := strconv.Atoi(record[14])
		p21Hijos, err := strconv.Atoi(record[15])
		p22GestionT, err := strconv.Atoi(record[16])
		p23FLapoyo, err := strconv.Atoi(record[17])
		p24Confami, err := strconv.Atoi(record[18])
		p25Estres, err := strconv.Atoi(record[19])
		p26Difiacad, err := strconv.Atoi(record[20])
		p27Tiempoest, err := strconv.Atoi(record[21])
		p28Promacu, err := strconv.Atoi(record[22])
		p29ApoDct, err := strconv.Atoi(record[23])
		p30CaliDoce, err := strconv.Atoi(record[24])
		p31MetoCalif, err := strconv.Atoi(record[25])
		p32PedaDoc, err := strconv.Atoi(record[26])
		p33GraMoti, err := strconv.Atoi(record[27])
		p34OpiAmig, err := strconv.Atoi(record[28])
		p35AcoSex, err := strconv.Atoi(record[29])
		p36PerfAcad, err := strconv.Atoi(record[30])
		p37CargAcad, err := strconv.Atoi(record[31])
		p38ValSem, err := strconv.Atoi(record[32])
		p39HabiPrac, err := strconv.Atoi(record[33])
		p40EntoUni, err := strconv.Atoi(record[34])
		p41Campus, err := strconv.Atoi(record[35])
		p42RecomUni, err := strconv.Atoi(record[36])
		p7_1Amigos, err := strconv.Atoi(record[37])
		p7_2FaciLabo, err := strconv.Atoi(record[38])
		p7_3Discipli, err := strconv.Atoi(record[39])
		p7_4Otro, err := strconv.Atoi(record[40])
		p7_5Pareja, err := strconv.Atoi(record[41])
		p7_6TradiFam, err := strconv.Atoi(record[42])
		p7_7VocaGusto, err := strconv.Atoi(record[43])
		p10_1DifiAcadem, err := strconv.Atoi(record[44])
		p10_2OrientUniversi, err := strconv.Atoi(record[45])
		p10_3FaltaIntCarrera, err := strconv.Atoi(record[46])
		p10_3FaltaMotiva, err := strconv.Atoi(record[47])
		p10_4NoHeInterrum, err := strconv.Atoi(record[48])
		p10_5ProbleSaludMen, err := strconv.Atoi(record[49])
		p10_6ProbleFinanci, err := strconv.Atoi(record[50])
		p10_7ProblePersonl, err := strconv.Atoi(record[51])
		p11_1Difieco, err := strconv.Atoi(record[52])
		p11_2Ubilab, err := strconv.Atoi(record[53])
		p11_3Cambciu, err := strconv.Atoi(record[54])
		p11_4Aplaseme, err := strconv.Atoi(record[55])
		p11_5Dififami, err := strconv.Atoi(record[56])
		p11_6BajoRenAca, err := strconv.Atoi(record[57])
		p11_7interpro, err := strconv.Atoi(record[58])
		p11_8Enfer, err := strconv.Atoi(record[59])
		p13_1Fallemadrepadre, err := strconv.Atoi(record[60])
		p13_2Ingrelabo, err := strconv.Atoi(record[61])
		p13_3Evenhabi, err := strconv.Atoi(record[62])
		p13_4Probpsico, err := strconv.Atoi(record[63])
		p13_5Perditrab, err := strconv.Atoi(record[64])
		p13_6Separapadres, err := strconv.Atoi(record[65])
		p13_7Sermadrepadre, err := strconv.Atoi(record[66])
		p13_8Ninguna, err := strconv.Atoi(record[67])
		p14_1CambiCarre, err := strconv.Atoi(record[68])
		p14_2CancSemest, err := strconv.Atoi(record[69])
		p14_3CancaMateria, err := strconv.Atoi(record[70])
		p14_4Ninguno, err := strconv.Atoi(record[71])
		p14_5RetirarPermanUniver, err := strconv.Atoi(record[72])
		p20_1Beca, err := strconv.Atoi(record[73])
		p20_2Credito, err := strconv.Atoi(record[74])
		p20_3Familiar, err := strconv.Atoi(record[75])
		p20_4TrabajoPropio, err := strconv.Atoi(record[76])

		encuesta := models.Encuesta{
			ID:                       uiD,
			P1Carrera:                p1Carrera,
			P2Semestre:               p2Semestre,
			P3Edad:                   p3Edad,
			P4Sisben:                 p4Sisben,
			P5Transporte:             p5Transporte,
			P6Orientacion:            p6Orientacion,
			P8Ideastu:                p8Ideastu,
			P9Perfil:                 p9Perfil,
			P12CambioU:               p12CambioU,
			P15Trabajo:               p15Trabajo,
			P16Ingresos:              p16Ingresos,
			P17Reside:                p17Reside,
			P18Foraneo:               p18Foraneo,
			P19Pcargo:                p19Pcargo,
			P21Hijos:                 p21Hijos,
			P22GestionT:              p22GestionT,
			P23FLapoyo:               p23FLapoyo,
			P24Confami:               p24Confami,
			P25Estres:                p25Estres,
			P26Difiacad:              p26Difiacad,
			P27Tiempoest:             p27Tiempoest,
			P28Promacu:               p28Promacu,
			P29ApoDct:                p29ApoDct,
			P30CaliDoce:              p30CaliDoce,
			P31MetoCalif:             p31MetoCalif,
			P32PedaDoc:               p32PedaDoc,
			P33GraMoti:               p33GraMoti,
			P34OpiAmig:               p34OpiAmig,
			P35AcoSex:                p35AcoSex,
			P36PerfAcad:              p36PerfAcad,
			P37CargAcad:              p37CargAcad,
			P38ValSem:                p38ValSem,
			P39HabiPrac:              p39HabiPrac,
			P40EntoUni:               p40EntoUni,
			P41Campus:                p41Campus,
			P42RecomUni:              p42RecomUni,
			P7_1Amigos:               p7_1Amigos,
			P7_2FaciLabo:             p7_2FaciLabo,
			P7_3Discipli:             p7_3Discipli,
			P7_4Otro:                 p7_4Otro,
			P7_5Pareja:               p7_5Pareja,
			P7_6TradiFam:             p7_6TradiFam,
			P7_7VocaGusto:            p7_7VocaGusto,
			P10_1DifiAcadem:          p10_1DifiAcadem,
			P10_2OrientUniversi:      p10_2OrientUniversi,
			P10_3FaltaIntCarrera:     p10_3FaltaIntCarrera,
			P10_3FaltaMotiva:         p10_3FaltaMotiva,
			P10_4NoHeInterrum:        p10_4NoHeInterrum,
			P10_5ProbleSaludMen:      p10_5ProbleSaludMen,
			P10_6ProbleFinanci:       p10_6ProbleFinanci,
			P10_7ProblePersonl:       p10_7ProblePersonl,
			P11_1Difieco:             p11_1Difieco,
			P11_2Ubilab:              p11_2Ubilab,
			P11_3Cambciu:             p11_3Cambciu,
			P11_4Aplaseme:            p11_4Aplaseme,
			P11_5Dififami:            p11_5Dififami,
			P11_6BajoRenAca:          p11_6BajoRenAca,
			P11_7interpro:            p11_7interpro,
			P11_8Enfer:               p11_8Enfer,
			P13_1Fallemadrepadre:     p13_1Fallemadrepadre,
			P13_2Ingrelabo:           p13_2Ingrelabo,
			P13_3Evenhabi:            p13_3Evenhabi,
			P13_4Probpsico:           p13_4Probpsico,
			P13_5Perditrab:           p13_5Perditrab,
			P13_6Separapadres:        p13_6Separapadres,
			P13_7Sermadrepadre:       p13_7Sermadrepadre,
			P13_8Ninguna:             p13_8Ninguna,
			P14_1CambiCarre:          p14_1CambiCarre,
			P14_2CancSemest:          p14_2CancSemest,
			P14_3CancaMateria:        p14_3CancaMateria,
			P14_4Ninguno:             p14_4Ninguno,
			P14_5RetirarPermanUniver: p14_5RetirarPermanUniver,
			P20_1Beca:                p20_1Beca,
			P20_2Credito:             p20_2Credito,
			P20_3Familiar:            p20_3Familiar,
			P20_4TrabajoPropio:       p20_4TrabajoPropio,
		}

		if err := config.DB.Create(&encuesta).Error; err != nil {
			panic(err)
		}
	}

}

func GetEncuestas(c *gin.Context) {
	encuestas := []models.Encuesta{}
	//se buscan los datos de la tabla
	config.DB.Find(&encuestas)
	//se responde al cliente la peticion ok y el dato tipo json
	c.IndentedJSON(http.StatusOK, &encuestas)
}
func PostEncuestas(c *gin.Context) {
	Read_csv()
	c.IndentedJSON(http.StatusCreated, gin.H{})
}
func GetEncuesta(c *gin.Context) {
	encuesta := models.Encuesta{}
	if err := config.DB.Where("ID = ?", c.Param("id")).First(&encuesta).Error; err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
	}
	c.IndentedJSON(http.StatusOK, &encuesta)
}
