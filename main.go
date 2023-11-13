package main

/*
С помощью горутин и каналов смоделировать работу пекарни:

Пекарня состоит из 2+1 этапов: торт выпекается, торт запаковывается, все торты смотрятся. Связь между этапами происходит через каналы, этапы выполняются с помощью горутин. Торт представляет из себя объект с полями типа int: BakedBy, BakeTime, PackedBy, PackTime Программа должна работать следующим образом:

Запускается N горутин, каждая из которых за время T1 = i +-t1 (i - номер рутины, t1 - выбранный вами параметр) создаёт объект тортика, заполняет поля BackedBy=i, BakeTime=T1 и отправляет его в канал

Существует пул из M горутин. когда в канал приходят торитики с предыдущего этапа, свободная рутина из пула начинает его упаковывать, то есть за время T2 = j+-t2 (j - номер рутины из пула, t2 - выбранный вами параметр, причём t2>=t1) выставляет тортику параметры PackedBy=j, PackTime = T2 и отправляет в канал

Мы ждём пока придут все тортики, или мы получим сигнал о завершении работы(тогда дожидаемся завершения работы текущих рутин), после чего просто выводим тортики в порядке, в котором они пришли к нам. При полном отрабатывании программы должно быть K тортиков.

Для данной задачи можно использовать различные наборы параметров, однако рекоммендуется попробовать следующие комбинации:

K = 10_000, N = 1, M = 1

K = 10_000, N = 8, M = 5

K = 10_000, N = 100, M = 50
*/

func main() {

}
