# Sera Yetiştirciliği

Seralarmızda akıllı ölçümlerle sebze yetiştiriciliği yapılacaktır. Seramızda yetiştirebileceğimiz bitki türleri en az 5 adet olacak şekilde önceden tanımlanacaktır. Akıllı sera ölçüm parametreleri olarak en az 3 adet parametre kullanılacaktır (ph , azot vs..). Bitkilerimizin saatlik olarak parametrelere yapmış olduğu katkı yada hasar (azotu artırıp azatlması gibi) takip edilecek buna göre akıllı sistem devreye girip yeniden serayı doğal şartlarına geri getirecektir ve gün sonunda günlük olarak üretilen sebzelerin miktarları kullanıcıya raporlanacaktır. 

## Yapılacaklar

- Var sayılan olarak en az 5 adet Bitki sisteme yazılım aşamasında tanımlanacak
- Yazılım başladığında kullanıcıya her bitkinin tüm özellikleri tek tek rapor edilecek
- Kullanıcı isterse konsoldan (ftm.Scan...) gireceği komutlarla çalışma anında default bitkilerin bilgilerini düzenleyebilecektir.
- Kullanıcı isterse konsoldan kendi serasının boyutlarını tanımlayabilecektir.
- Kullanıcı başlat komutu verdiğinde serada ekim işlemleri yapılacak ve saatlik (her saat sistem için 1 saniye olabilir) ölçümlere göre akıllı sistem devreye girip parametrik düzenlemeler yapılacaktır ve sonuçları adım adım ekranda yazılacaktır.
- Günlük olarak 24 saate bir ( 24 sn) toplam üretilen ve hasat edilen bitkilerin topnajları kullanıcıya iletilecektir.

## Çalışma Prensibi

- Seraların içerisinde ekim yapılacak sıralara Random olarak olarak bitkiler yerşeltirilecek. Her sıranın değerleri saatlik olarak ( 1 sn ) sistemde gösterilecek. Yapılan ölçümler neticesinde akıllı sistem devreye alınıp sıra sıra değerler ph azot vs.. ekleme çıkaratma yapılarak default hallerine geri gelecek.
- Her sıraya Random olarak seçilen sadece bir bitki türü yerleştirilecek ve o sıranın ihtiyacı olan parametrik veri o bitkiden alınacaktır.


## Notlar

- Sebze tanımlamalarında Struct kullanılacak.
- Bitkilerin özellerikleri tanımlanırken "methotlar" kullanılacak (foknsiyonlar değil).
- Bitkiler in-memory map olarak hafızada tutlacak ( map kullanılacak ).
- Log akışları sırasında time.Sleep ile log akışı yavaşlatılırsa daha gerçekci olacaktır.
- Yazılım ilk başladığında nasıl çalıştığını kullanıcın neler yapabileceğini anlatan bir çıktıyı ekranda göstermelidir.
