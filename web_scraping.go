package main

import (
	"fmt"
	"math/rand"
	"time"
)

type SeraS struct {
	m3                          uint16
	plantsInsıde                string
	co2                         uint16
	avarageWaterRequariment     uint16
	avarageWattRequariment      uint16
	minTemperature              int16
	maxTemparute                int16
	maxGraundHumiduty           uint16
	minGraundHumiduty           uint16
	risingGroundResult          uint16
	pricePerOneDegreeElectritcy uint16
	waterBillPerM3              uint16
	cropPerDay                  uint16
	totalCrpo                   uint16
}

func init() {
	fmt.Print("\n\n")
	fmt.Println("        ******************************")
	fmt.Println("             Akıllı sera sistemi      ")

	fmt.Println("        ******************************")
	fmt.Print("\n\n")
	fmt.Println("-->Akıllı sera sistemi seranıza ektiğiniz bitkilerin gereksinimlerini kaydeder")
	fmt.Println("-->Seranızdaki günlük co2 , sıcaklık ,  ve topraktaki nem oranındaki değişimleri hesaplayarak bu değerleri optimum değerlere çevirir.")
	fmt.Println("-->Sistemimze seranızın boyutlarını düzenleytebilirsiniz.")
	fmt.Println("-->Düzenlemelerinizi yaptıktan sonra sistem devreye girer ve size gün başı  raporlama işlemi yapar.")
	fmt.Println("-->Bu raporda akıllı sistemin seranızda yaptığı değişiklikler yer alır")
	fmt.Println("-->Sistemimiz gün sonunda toplam üretilen ve hasat edilen bitkilerin topnajları ve toplam faturayı rapor edecektir")
	fmt.Print("\n\n")

}

func main() {

	var (
		i                        uint16
		k                        uint16
		plantVariety             uint8
		greenHouse               SeraS
		height                   uint16
		m2                       uint16
		numberOfDay              uint16
		greenHouseDayLong        []uint16
		plantVarietyTwo          []string
		electricicityBill        uint16
		totalCrop                uint16
		plantt                   string
		airCo2Ppm                int16
		temeratureOfDay          int16
		changeOfTemperature      int16
		plantDay                 uint16
		risingTemparatureVar     int16
		decraisingTemparature    int16
		increasingGroundHumidity uint16
		decreasingGroundHumidity uint16
		day                      uint16
		totalWaterBill           uint16
		dailyWaterBill           uint16
		dailyElectricityBill     uint16
		totalDayCrop             uint16
		totalHumidity            uint16
	)
	var electricityBillsCount [5]uint16 = [5]uint16{0, 0, 0, 0, 0}
	var waterBillsCount [5]uint16 = [5]uint16{0, 0, 0, 0, 0}
	var kgsOfCrops [5]uint16 = [5]uint16{0, 0, 0, 0, 0}
	temeratureOfDay = 20
	greenHouse.pricePerOneDegreeElectritcy = 50
	greenHouse.waterBillPerM3 = 2

	fmt.Println("sera yükseklik belirleyin")
	fmt.Println("lütfen bir tam sayi yaziniz")
	fmt.Scan(&height)

	fmt.Println("sera m2 belirleyin")
	fmt.Println("lütfen bir tam sayi yaziniz")
	fmt.Scan(&m2)

	greenHouse.m3 = height * m2
	fmt.Println("Sera yüksekliği = ", greenHouse.m3, " m3")

	fmt.Println("Sistemde özellikleri bulunan bitkliler = marul , havuç , domates , soğan , patates , salatalık")
	fmt.Println("sera da yetiştireceğiniz bitki çeşit sayısını belirleyiniz")
	fmt.Scan(&plantVariety)
	fmt.Println("Bitki çesidi saysı =", plantVariety)
	fmt.Println("Serada yetiştireceğiniz bitkileri belirleyiniz")
	for k = 0; k < uint16(plantVariety); k++ { // kullanıcıdan bitkileri alıyoruz

		fmt.Scan(&plantt)

		plantVarietyTwo = append(plantVarietyTwo, plantt)

	}

	for t := 0; t < len(plantVarietyTwo); t++ { // bitkilerin günlerini alıyoruz
		fmt.Println(t, ". Bitki =", plantVarietyTwo[t])
		dayOfGreenHouse := greenHouse.bitkiOzelligi(plantVarietyTwo[t])
		greenHouseDayLong = append(greenHouseDayLong, dayOfGreenHouse)
		fmt.Println(t, ". Bitkinin gun sayısı", dayOfGreenHouse)

	}

	day = yuksekGunSayisiHesapla(greenHouseDayLong) // en uzun güne sahip bitkinin gününü alıyoruz
	for i = 0; i < uint16(day); i++ {

		time.Sleep(2 * time.Second)
		numberOfDay++
		fmt.Printf("***********************\n")
		fmt.Printf("|    Gun sayisi = %d   |\n", i+1)

		rand.Seed(time.Now().UnixNano())

		changeOfTemperature = int16(rand.Intn(5 - (-5)))
		totalHumidity = uint16(rand.Intn(15 - 0))
		temeratureOfDay += changeOfTemperature
		fmt.Printf("|    Hava = %d derece |", temeratureOfDay)
		fmt.Print("\n")
		fmt.Println("***********************")
		for v := 0; v < len(plantVarietyTwo); v++ {

			plantDay = greenHouse.bitkiOzelligi(plantVarietyTwo[v])
			if i >= plantDay {
				continue
			}
			fmt.Printf("\n\n\n")
			fmt.Printf("%s serası için değerler", plantVarietyTwo[v])
			fmt.Println("")

			if totalHumidity < greenHouse.minGraundHumiduty {
				increasingGroundHumidity = greenHouse.minGraundHumiduty - totalHumidity
				fmt.Printf("-->Toprak nemi yuzde ,%d kadar arttırıldı\n", increasingGroundHumidity)
				dailyWaterBill = greenHouse.m3 * greenHouse.waterBillPerM3 * increasingGroundHumidity
				totalWaterBill = dailyWaterBill + totalWaterBill
				waterBillsCount[v] = totalWaterBill
				fmt.Printf("-->Günlük su faturası = %d\n", dailyWaterBill)
				fmt.Printf("-->Toplam Su faturası %d\n", waterBillsCount[(v)])

			} else if totalHumidity > greenHouse.maxGraundHumiduty {
				decreasingGroundHumidity = totalHumidity - uint16(greenHouse.maxGraundHumiduty)

				fmt.Printf("-->Toprak nemi yuzde ,%d kadar azaltıldı", decreasingGroundHumidity)
				fmt.Println()

				dailyWaterBill = greenHouse.m3 * greenHouse.waterBillPerM3 * decreasingGroundHumidity

				totalWaterBill = dailyWaterBill + totalWaterBill
				waterBillsCount[v] = totalWaterBill
				fmt.Printf("-->Günlük su faturası = %d\n", dailyWaterBill)
				fmt.Printf("-->Toplam Su faturası %d\n", waterBillsCount[v])

			} else {

				fmt.Println("-->Nem normal durumda")

			}

			if temeratureOfDay < greenHouse.minTemperature { // ortamdaki co2 seviyesini  sıcaklığa göre belirledik

				risingTemparatureVar = greenHouse.minTemperature - temeratureOfDay
				fmt.Printf("-->Hava sıcaklıgı olması gerekenden %d kadar düşük\n", risingTemparatureVar)
				airCo2Ppm = risingTemparatureVar * 45
				fmt.Printf("-->Sıcaklık %d°  arttırıldı\n", risingTemparatureVar)

				fmt.Printf("-->CO2 %d  kadar arttırıldı\n", airCo2Ppm)

				dailyElectricityBill = greenHouse.m3 * greenHouse.pricePerOneDegreeElectritcy * uint16(risingTemparatureVar)
				fmt.Printf("-->Gunluk elektrik faturası = %d\n", dailyElectricityBill)
				electricicityBill = dailyElectricityBill + electricicityBill
				electricityBillsCount[v] = electricicityBill

				fmt.Printf("\n\n\n")

			} else if temeratureOfDay > greenHouse.maxTemparute {
				decraisingTemparature = temeratureOfDay - greenHouse.maxTemparute
				fmt.Printf("-->Hava sıcaklıgı olması gerekenden %d kadar yüksek\n", decraisingTemparature)

				airCo2Ppm = decraisingTemparature * 45
				fmt.Printf("-->Sıcaklık %d°  azaltıldı\n", decraisingTemparature)

				fmt.Printf("-->CO2 %d  kadar azaltıldı\n", airCo2Ppm)

				dailyElectricityBill = greenHouse.m3 * greenHouse.pricePerOneDegreeElectritcy * uint16(decraisingTemparature)
				electricityBillsCount[v] = electricicityBill
				fmt.Printf("-->Gunluk elektrik faturası = %d\n", dailyElectricityBill)

			} else {

				fmt.Println("-->Sıcaklık ve CO2 normal durumda")

			}
			fmt.Printf("-->Toplam Elektrik Faturası = %d tl\n", electricityBillsCount[v])
			totalDayCrop = greenHouse.cropPerDay * m2
			totalCrop = totalDayCrop + totalCrop

			kgsOfCrops[v] = totalCrop

			if (i+1)%4 == 0 {

				fmt.Printf("-->şuana kadar ki toplam %d mahsulu", kgsOfCrops[v])
				fmt.Printf("\n\n\n")
			}

		}
		temeratureOfDay = 20
	}
	fmt.Printf("\n\n\n")
	fmt.Printf("\n\n\n")
	for v := 0; v < len(plantVarietyTwo); v++ {
		fmt.Printf("-->%s icin Odenmesi gereken toplam elektrik farurası %d tl\n", plantVarietyTwo[v], electricityBillsCount[(v)])
		fmt.Printf("-->%s icin Odenmesi gereken toplam su farurası %d tl\n", plantVarietyTwo[v], waterBillsCount[(v)])
		fmt.Printf("-->%s icin toplam mahsul %d kg\n", plantVarietyTwo[v], kgsOfCrops[(v)])
		fmt.Printf("\n\n\n")
	}
}
func (a *SeraS) bitkiOzelligi(bitki string) uint16 {
	var gun uint16
	var bitkiGunu map[string]uint16
	bitkiGunu = make(map[string]uint16)

	bitkiGunu["marul"] = 90
	bitkiGunu["havuç"] = 120
	bitkiGunu["domates"] = 100
	bitkiGunu["patates"] = 90
	bitkiGunu["soğan"] = 85
	bitkiGunu["salatalık"] = 110

	for k, v := range bitkiGunu {
		switch {

		case k == "marul" && bitki == "marul":
			a.co2 = 1200
			a.avarageWaterRequariment = 3
			a.avarageWattRequariment = 5
			gun = v
			a.minTemperature = 22
			a.maxTemparute = 27
			a.minGraundHumiduty = 8
			a.maxGraundHumiduty = 11
			a.cropPerDay = 10
			a.totalCrpo = 0

		case k == "havuç" && bitki == "havuç":
			a.co2 = 1100
			a.avarageWaterRequariment = 2
			a.avarageWattRequariment = 8
			gun = v
			a.minTemperature = 23
			a.maxTemparute = 26
			a.minGraundHumiduty = 9
			a.maxGraundHumiduty = 14
			a.cropPerDay = 8
			a.totalCrpo = 0

		case k == "domates" && bitki == "domates":
			a.co2 = 1500
			a.avarageWaterRequariment = 1
			a.avarageWattRequariment = 3
			gun = v
			a.minTemperature = 24
			a.maxTemparute = 29
			a.minGraundHumiduty = 10
			a.maxGraundHumiduty = 14
			a.cropPerDay = 6
			a.totalCrpo = 0

		case k == "patates" && bitki == "patates":
			a.co2 = 900
			a.avarageWaterRequariment = 1
			a.avarageWattRequariment = 4
			gun = v
			a.minTemperature = 21
			a.maxTemparute = 28
			a.minGraundHumiduty = 11
			a.maxGraundHumiduty = 13
			a.cropPerDay = 3
			a.totalCrpo = 0

		case k == "soğan" && bitki == "soğan":
			a.co2 = 1000
			a.avarageWaterRequariment = 3
			a.avarageWattRequariment = 7
			gun = v
			a.minTemperature = 21
			a.maxTemparute = 27
			a.minGraundHumiduty = 9
			a.maxGraundHumiduty = 11
			a.cropPerDay = 12
			a.totalCrpo = 0

		case k == "salatalık" && bitki == "salatalık":
			a.co2 = 1100
			a.avarageWaterRequariment = 2
			a.avarageWattRequariment = 5
			gun = v
			a.minTemperature = 21
			a.maxTemparute = 23
			a.minGraundHumiduty = 5
			a.maxGraundHumiduty = 10
			a.cropPerDay = 14
			a.totalCrpo = 0

		}
	}
	return gun
}
func yuksekGunSayisiHesapla(slice []uint16) uint16 {
	var (
		a uint16
		i uint16
	)
	for i = 0; i < uint16(len(slice)); i++ {

		if slice[i] > a {
			a = slice[i]
		}

	}

	return a

}
