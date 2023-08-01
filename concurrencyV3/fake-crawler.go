package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)


func goColly(urlCh <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range(urlCh) {
		// Go Colly Aqui
		sleepTime := rand.Intn(1) + 1
		time.Sleep(time.Duration(sleepTime) * time.Second)
		fmt.Printf("Crawled: %s\nTimetaken: %d\n\n", url, sleepTime)
	}
}

func Crawl(urls []string) {
	var maxConcurrentRequests = 5
	var wg sync.WaitGroup

	urlCh := make(chan string)

	for i := 0; i < maxConcurrentRequests; i++ {
		wg.Add(1)
		go goColly(urlCh, &wg)
	}

	for _, url := range urls {
		urlCh <- fmt.Sprintf("https://amazon.com.br/dp/%s", url)
	}
	close(urlCh)

	wg.Wait()
}

func main() {
	var urls = []string{"https://www.sample.org/?destruction=walk&seashore=shelf", "https://www.sample.net/?quicksand=drawer&teeth=earthquake#digestion", "http://sample.org/?rose=mouth&account=mint", "https://action.sample.com/", "https://www.sample.org/store", "https://sample.net/ice#mouth", "http://sample.net/?grip=office", "https://www.wrench.sample.com/", "https://sample.info/expansion#front", "https://www.sample.org/thrill.html", "http://www.sample.edu/?toy=umbrella&offer=chess", "http://sample.org/?hand=tray&daughter=punishment", "https://www.sample.info/?queen=view&voice=pigs", "https://www.sample.info/hole.php", "https://www.sample.com/part", "http://sample.info/wire", "https://drain.sample.net/", "https://www.sample.edu/?gold=measure&car=transport", "https://www.sample.edu/partner.aspx", "http://sample.edu/purpose#truck", "http://sample.org/?sack=wilderness&earth=wind", "https://sample.com/?back=apparel&business=distance#addition", "http://sample.com/veil.html", "https://sample.org/?desk=wire&ground=paper", "https://www.legs.sample.net/foot/committee?connection=worm#flavor", "https://channel.sample.com/fold/action.html", "https://sample.edu/skate", "http://sample.info/#heat", "https://sample.net/?reaction=bait&punishment=soup#day", "http://www.sample.edu/?rain=pancake&neck=scissors", "http://sample.com/?pleasure=value&payment=scene#cherry", "http://sample.org/library?profit=color", "http://www.sample.com/birth", "https://www.sample.net/finger#van", "https://sample.com/screw", "https://sample.org/?seed=land", "https://www.sample.edu/#pen", "https://scene.sample.com/", "http://www.sample.com/milk.aspx", "http://www.sample.com/sail#bottle", "http://www.sample.org/amusement", "https://sample.info/basketball#slope", "https://sample.com/water?fly=thunder", "http://sample.info/?sort=jeans&wish=money#sister", "http://www.sample.edu/rain", "http://sample.com/?building=amount&pull=texture", "https://www.sample.com/airplane?mailbox=laborer", "https://www.sample.org/#mountain", "http://sample.edu/?debt=talk", "https://sample.info/?chance=health&rub=nation#fire", "https://www.sample.info/?cover=bike&pizzas=profit", "http://www.sample.com/", "http://www.clover.sample.edu/", "https://www.planes.sample.info/", "https://www.sample.net/?recess=force", "http://sample.com/condition.aspx", "https://sample.info/glass#beds", "http://sample.info/#cobweb", "http://www.sample.net/cream#pot", "https://sample.net/?beef=route&pancake=surprise#bed", "http://arm.sample.com/quill/bit.html", "https://sample.org/?taste=scale&noise=potato", "https://sample.com/leather#religion", "http://sample.org/frogs.html", "https://sample.net/spot#books", "https://nose.sample.info/snow/rabbits.aspx", "http://www.sample.edu/?birth=rod&language=impulse#color", "http://sample.org/?rule=books&stage=mask", "https://sample.com/?gun=fall&fuel=nose", "http://sample.info/meeting#seashore", "https://www.sample.com/profit?change=eggs", "http://sample.org/?picture=collar&rainstorm=drain", "http://sample.com/support?skirt=hill", "https://sample.net/?smell=brass", "https://www.sample.com/number", "http://sample.info/?blade=hot&sack=rate", "https://sample.edu/silk.html", "https://sample.info/?country=cent", "https://www.sample.edu/silk", "https://www.sample.com/fire", "https://www.sample.org/cherry?relation=drawer", "https://www.range.sample.edu/nation/market?chickens=loaf#meal", "https://sample.info/food.php", "http://sample.org/?ice=act&power=committee", "http://www.sample.org/?truck=sound&taste=value", "https://www.sample.info/?noise=middle&driving=visitor", "http://war.sample.net/", "https://sample.net/cover.aspx", "http://sample.info/visitor", "https://marble.sample.net/", "http://sample.info/?stop=care&reason=daughter", "http://sample.org/?company=ear&sweater=train#bomb", "http://sample.com/property?invention=trip", "https://www.sample.net/zinc#sneeze", "https://www.rainstorm.sample.org/coal/cellar.php", "http://sample.edu/", "https://snow.sample.net/"}
	Crawl(urls)
}
