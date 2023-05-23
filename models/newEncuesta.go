package models

import (
	"gorm.io/gorm"
)

type NewEncuesta struct {
	gorm.Model
	ID                       uint `gorm:"column:id;primaryKey;unique;autoIncrement:false"`
	P1Carrera                int  `gorm:"column:1Carrera"`
	P2Semestre               int  `gorm:"column:2Semestre"`
	P3Edad                   int  `gorm:"column:3Edad"`
	P4Sisben                 int  `gorm:"column:4Sisben"`
	P5Transporte             int  `gorm:"column:5Transporte"`
	P6Orientacion            int  `gorm:"column:6Orientacion"`
	P8Ideastu                int  `gorm:"column:8Ideastu"`
	P9Perfil                 int  `gorm:"column:9Perfil"`
	P12CambioU               int  `gorm:"column:12CambioU"`
	P15Trabajo               int  `gorm:"column:15Trabajo"`
	P16Ingresos              int  `gorm:"column:16Ingresos"`
	P17Reside                int  `gorm:"column:17Reside"`
	P18Foraneo               int  `gorm:"column:18Foraneo"`
	P19Pcargo                int  `gorm:"column:19Pcargo"`
	P21Hijos                 int  `gorm:"column:21Hijos"`
	P22GestionT              int  `gorm:"column:22GestionT"`
	P23FLapoyo               int  `gorm:"column:23FLapoyo"`
	P24Confami               int  `gorm:"column:24Confami"`
	P25Estres                int  `gorm:"column:25Estres"`
	P26Difiacad              int  `gorm:"column:26Difiacad"`
	P27Tiempoest             int  `gorm:"column:27Tiempoest"`
	P28Promacu               int  `gorm:"column:28Promacu"`
	P29ApoDct                int  `gorm:"column:29ApoDct"`
	P30CaliDoce              int  `gorm:"column:30CaliDoce"`
	P31MetoCalif             int  `gorm:"column:31MetoCalif"`
	P32PedaDoc               int  `gorm:"column:32PedaDoc"`
	P33GraMoti               int  `gorm:"column:33GraMoti"`
	P34OpiAmig               int  `gorm:"column:34OpiAmig"`
	P35AcoSex                int  `gorm:"column:35AcoSex"`
	P36PerfAcad              int  `gorm:"column:36PerfAcad"`
	P37CargAcad              int  `gorm:"column:37CargAcad"`
	P38ValSem                int  `gorm:"column:38ValSem"`
	P39HabiPrac              int  `gorm:"column:39HabiPrac"`
	P40EntoUni               int  `gorm:"column:40EntoUni"`
	P41Campus                int  `gorm:"column:41Campus"`
	P42RecomUni              int  `gorm:"column:42RecomUni"`
	P7_1Amigos               int  `gorm:"column:7_1Amigos"`
	P7_2FaciLabo             int  `gorm:"column:7_2FaciLabo"`
	P7_3Discipli             int  `gorm:"column:7_3Discipli"`
	P7_4Otro                 int  `gorm:"column:7_4Otro"`
	P7_5Pareja               int  `gorm:"column:7_5Pareja"`
	P7_6TradiFam             int  `gorm:"column:7_6TradiFam"`
	P7_7VocaGusto            int  `gorm:"column:7_7VocaGusto"`
	P10_1DifiAcadem          int  `gorm:"column:10_1DifiAcadem"`
	P10_2OrientUniversi      int  `gorm:"column:10_2OrientUniversi"`
	P10_3FaltaIntCarrera     int  `gorm:"column:10_3FaltaIntCarrera"`
	P10_3FaltaMotiva         int  `gorm:"column:10_3FaltaMotiva"`
	P10_4NoHeInterrum        int  `gorm:"column:10_4NoHeInterrum"`
	P10_5ProbleSaludMen      int  `gorm:"column:10_5ProbleSaludMen"`
	P10_6ProbleFinanci       int  `gorm:"column:10_6ProbleFinanci"`
	P10_7ProblePersonl       int  `gorm:"column:10_7ProblePersonl"`
	P11_1Difieco             int  `gorm:"column:11_1Difieco"`
	P11_2Ubilab              int  `gorm:"column:11_2Ubilab"`
	P11_3Cambciu             int  `gorm:"column:11_3Cambciu"`
	P11_4Aplaseme            int  `gorm:"column:11_4Aplaseme"`
	P11_5Dififami            int  `gorm:"column:11_5Dififami"`
	P11_6BajoRenAca          int  `gorm:"column:11_6BajoRenAca"`
	P11_7interpro            int  `gorm:"column:11_7interpro"`
	P11_8Enfer               int  `gorm:"column:11_8Enfer"`
	P13_1Fallemadrepadre     int  `gorm:"column:13_1Fallemadrepadre"`
	P13_2Ingrelabo           int  `gorm:"column:13_2Ingrelabo"`
	P13_3Evenhabi            int  `gorm:"column:13_3Evenhabi"`
	P13_4Probpsico           int  `gorm:"column:13_4Probpsico"`
	P13_5Perditrab           int  `gorm:"column:13_5Perditrab"`
	P13_6Separapadres        int  `gorm:"column:13_6Separapadres"`
	P13_7Sermadrepadre       int  `gorm:"column:13_7Sermadrepadre"`
	P13_8Ninguna             int  `gorm:"column:13_8Ninguna"`
	P14_1CambiCarre          int  `gorm:"column:14_1CambiCarre"`
	P14_2CancSemest          int  `gorm:"column:14_2CancSemest"`
	P14_3CancaMateria        int  `gorm:"column:14_3CancaMateria"`
	P14_4Ninguno             int  `gorm:"column:14_4Ninguno"`
	P14_5RetirarPermanUniver int  `gorm:"column:14_5RetirarPermanUniver"`
	P20_1Beca                int  `gorm:"column:20_1Beca"`
	P20_2Credito             int  `gorm:"column:20_2Credito"`
	P20_3Familiar            int  `gorm:"column:20_3Familiar"`
	P20_4TrabajoPropio       int  `gorm:"column:20_4TrabajoPropio"`
}
