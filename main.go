// How I'll approach this task
// 1. Create a simple http server	-	done
// 2. Make models for participant and the meeting	-	done
// 3. Apply the routes	-	done
// 4. Make handlers for each request	-	done 2/4 T_T
// 5. Go for detail structuring as given in the doc	-	couldn't do T_T
// 6. Document the code in a good way so that it is readable :)	-	done

package main

import (
	"log"
	"net/http"

	"./handlers" // handlers package imported to apply handler to http endpoints
)

func main() {
	http.HandleFunc("/meetings", handlers.ScheduleMeeting) // endpoint for POST on /meetings
	http.HandleFunc("/meetings/", handlers.GetMeetingByID) // endpoint for GET on /meetings/:id

	// didn't get time to implement the other two routes

	log.Fatal(http.ListenAndServe(":8090", nil))
}
