/* Copyright (C) 2015 by Alexandru Cojocaru */

/* This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>. */

package logparse

import (
	"fmt"
	"log"
)

func ExampleCommon() {
	l, err := Common(`:: - xojoc [10/Feb/2015:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326`)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(l)
	fmt.Println(l.Host)
	fmt.Println(l.User)
	fmt.Println(l.Time.Format("02/Jan/2006"))
	fmt.Println(l.Request.URL)
	fmt.Println(l.Status)
	fmt.Println(l.Bytes)
	//Output::: - xojoc [10/Feb/2015:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326 - -
	//::
	//xojoc
	//10/Feb/2015
	///apache_pb.gif
	//200
	//2326
}

func ExampleCombined() {
	l, err := Combined(`:: - xojoc [10/Feb/2015:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326 "http://xojoc.pw" "Mozilla/5.0 (X11; Linux i686; rv:38.0) Gecko/20100101 Firefox/38.0"`)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(l)
	fmt.Println(l.Host)
	fmt.Println(l.User)
	fmt.Println(l.Time.Format("02/Jan/2006"))
	fmt.Println(l.Request.URL)
	fmt.Println(l.Status)
	fmt.Println(l.Bytes)
	fmt.Println(l.Referer)
	fmt.Println(l.UserAgent.Name)
	fmt.Println(l.UserAgent.OS)
	//Output::: - xojoc [10/Feb/2015:13:55:36 -0700] "GET /apache_pb.gif HTTP/1.0" 200 2326 "http://xojoc.pw" "Mozilla/5.0 (X11; Linux i686; rv:38.0) Gecko/20100101 Firefox/38.0"
	//::
	//xojoc
	//10/Feb/2015
	///apache_pb.gif
	//200
	//2326
	//http://xojoc.pw
	//Firefox
	//GNU/Linux
}
