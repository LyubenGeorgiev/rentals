package main

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
	"time"
)

var (
	rentals = [][]byte{
		[]byte(`{"id":1,"name":"'Abaco' VW Bay Window: Westfalia Pop-top","description":"ultrices consectetur torquent posuere phasellus urna faucibus convallis fusce sem felis malesuada luctus diam hendrerit fermentum ante nisl potenti nam laoreet netus est erat mi","type":"camper-van","make":"Volkswagen","model":"Bay Window","year":1978,"length":15,"sleeps":4,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1528586451/p/rentals/4447/images/yd7txtw4hnkjvklg8edg.jpg","price":{"day":16900},"location":{"city":"Costa Mesa","state":"CA","zip":"92627","country":"US","lat":33.64,"lng":-117.93},"user":{"id":1,"first_name":"John","last_name":"Smith"}}`),
		[]byte(`{"id":2,"name":"Maupin: Vanagon Camper","description":"fermentum nullam congue arcu sollicitudin lacus suspendisse nibh semper cursus sapien quis feugiat maecenas nec turpis viverra gravida risus phasellus tortor cras gravida varius scelerisque","type":"camper-van","make":"Volkswagen","model":"Vanagon Camper","year":1989,"length":15,"sleeps":4,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1498568017/p/rentals/11368/images/gmtye6p2eq61v0g7f7e7.jpg","price":{"day":15000},"location":{"city":"Portland","state":"OR","zip":"97202","country":"US","lat":45.51,"lng":-122.68},"user":{"id":2,"first_name":"Jane","last_name":"Doe"}}`),
		[]byte(`{"id":3,"name":"1984 Volkswagen Westfalia","description":"urna iaculis sed ut porttitor mollis ante cubilia ad felis duis varius mollis nascetur metus faucibus ligula ultricies in faucibus morbi imperdiet auctor morbi torquent","type":"camper-van","make":"Volkswagen","model":"Westfalia","year":1984,"length":16,"sleeps":4,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1504395813/p/rentals/21399/images/nxtwdubpapgpmuc65pd1.jpg","price":{"day":18000},"location":{"city":"San Diego","state":"CA","zip":"92037","country":"US","lat":32.83,"lng":-117.28},"user":{"id":3,"first_name":"Barry","last_name":"Martin"}}`),
		[]byte(`{"id":4,"name":"Sm. #1 (Sleeps 2) - Check Dates for Price","description":"aliquet sit placerat libero viverra hendrerit ridiculus etiam pulvinar faucibus tempor magnis litora neque varius volutpat mollis class laoreet quisque montes cubilia leo aliquet litora","type":"camper-van","make":"Ford","model":"Transit 350","year":2016,"length":19,"sleeps":2,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1508688886/p/rentals/25403/images/jkqxknddnuq6fvmyatke.jpg","price":{"day":8900},"location":{"city":"Salt Lake City","state":"UT","zip":"84104","country":"US","lat":40.73,"lng":-111.92},"user":{"id":4,"first_name":"Todd","last_name":"Edison"}}`),
		[]byte(`{"id":5,"name":"Stardust2005Mercedes-BenzSprinter","description":"pretium sit in quis semper ligula sed sagittis molestie et vehicula cursus ullamcorper est euismod diam massa sem cum lorem cursus euismod vivamus urna leo","type":"camper-van","make":"Mercedes-Benz","model":"Sprinter","year":2005,"length":20,"sleeps":4,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1521261348/p/rentals/40129/images/wn0tx6meifqtrnwjmeoq.jpg","price":{"day":8000},"location":{"city":"San Diego","state":"CA","zip":"92109","country":"US","lat":32.8,"lng":-117.24},"user":{"id":5,"first_name":"Ben","last_name":"Reynard"}}`),
		[]byte(`{"id":6,"name":"2003 Winnebago Eurovan Camper Eurovan Camper","description":"eros tellus quisque tellus parturient elit varius maecenas justo aliquet metus neque sociis interdum commodo curae class leo massa cursus auctor nisl ante semper habitant","type":"camper-van","make":"Winnebago Eurovan Camper","model":"Eurovan Camper","year":2003,"length":17,"sleeps":4,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1523649590/p/rentals/46190/images/elinlzv6fpnrktik4wqh.jpg","price":{"day":13000},"location":{"city":"Charleston","state":"SC","zip":"29412","country":"US","lat":32.69,"lng":-79.96},"user":{"id":1,"first_name":"John","last_name":"Smith"}}`),
		[]byte(`{"id":7,"name":"2002 Volkswagen Eurovan Weekender Westfalia","description":"purus neque pellentesque potenti posuere molestie vivamus urna faucibus class justo porta litora turpis cubilia sit class torquent ullamcorper netus ut sapien libero consequat quisque","type":"camper-van","make":"VW","model":"Eurovan Weekender Westfalia","year":2002,"length":0,"sleeps":4,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1526614056/p/rentals/52210/images/nou2lx0h0dsjzbqeotuf.jpg","price":{"day":15000},"location":{"city":"Rancho Mission Viejo","state":"CA","zip":"","country":"US","lat":33.53,"lng":-117.63},"user":{"id":2,"first_name":"Jane","last_name":"Doe"}}`),
		[]byte(`{"id":8,"name":"2017 Transit Adventure Van","description":"commodo congue platea magnis montes feugiat lorem metus nullam ante convallis nulla dolor mauris praesent mus ante varius per hac sed metus auctor ultricies diam","type":"camper-van","make":"Ford","model":"Sacramento","year":2017,"length":20,"sleeps":2,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1562023338/p/rentals/119031/images/wchguimw6h3u9oonba9b.jpg","price":{"day":16500},"location":{"city":"Sacramento","state":"CA","zip":"95811","country":"US","lat":38.57,"lng":-121.49},"user":{"id":3,"first_name":"Barry","last_name":"Martin"}}`),
		[]byte(`{"id":9,"name":"Maui \"Alani\" camping car SUBARU IMPREZA 4WD  -Cold AC.","description":"fermentum torquent hac id tortor conubia litora proin sociosqu congue elit ridiculus fames velit viverra faucibus eleifend sagittis etiam aptent sociosqu taciti metus iaculis quam","type":"camper-van","make":"SUBARU IMPREZA 4WD","model":"SUBARU IMPREZA 4WD","year":2003,"length":13,"sleeps":2,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1538027810/p/rentals/82458/images/bphrohl2r4wxc8wg3v11.jpg","price":{"day":5900},"location":{"city":"Kahului","state":"HI","zip":"96732","country":"US","lat":20.88,"lng":-156.45},"user":{"id":4,"first_name":"Todd","last_name":"Edison"}}`),
		[]byte(`{"id":10,"name":"Betty!    1987 Volkswagen Westfalia Poptop Manual with kitchen!","description":"mollis curabitur cum convallis sagittis feugiat lectus ligula porta libero parturient maecenas cum facilisis ridiculus mauris ut est scelerisque tincidunt quisque hac lectus mus dapibus","type":"camper-van","make":"Volkswagen","model":"Westfalia","year":1987,"length":15,"sleeps":4,"primary_image_url":"https://res.cloudinary.com/outdoorsy/image/upload/v1535836865/p/rentals/91133/images/blijuwlisflua72ay1p2.jpg","price":{"day":25000},"location":{"city":"Missoula ","state":"MT","zip":"59808","country":"US","lat":46.92,"lng":-114.09},"user":{"id":5,"first_name":"Ben","last_name":"Reynard"}}`),
	}
)

func TestServerIntegration(t *testing.T) {
	// Wait a bit to allow the server to start
	time.Sleep(2 * time.Second)

	t.Log("Integration testing...\n")
	for i := 1; i <= len(rentals); i++ {
		// Make a request to the server
		resp, err := http.Get(fmt.Sprintf("http://app:8080/rentals/%d", i))
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("Error reading body: %v", err)
		}

		if !reflect.DeepEqual(rentals[i-1], bodyBytes) {
			t.Fatalf("Response is different from expected. Expected: %v\nActual: %v\n", rentals[i-1], bodyBytes)
		}

		resp.Body.Close()
	}
}
